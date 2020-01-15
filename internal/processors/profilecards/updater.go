package profilecards

import (
	"context"
	"errors"
	"sync"
	"time"

	"cirello.io/oversight"
	"github.com/cskr/pubsub"
	"github.com/faceit/go-steam"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	"github.com/13k/night-stalker/models"
)

const (
	processorName  = "profile_cards"
	cardExpiration = 1 * time.Hour
)

var (
	ErrInvalidProfileCardResponse = errors.New("Invalid profile card response")
)

type UpdaterOptions struct {
	Logger          *nslog.Logger
	PoolSize        int
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Updater)(nil)

type Updater struct {
	options           *UpdaterOptions
	ctx               context.Context
	log               *nslog.Logger
	db                *gorm.DB
	steam             *steam.Client
	dota              *dota2.Dota2
	workerPool        *ants.Pool
	activeReqsMtx     sync.Mutex
	activeReqs        map[uint32]bool
	bus               *pubsub.PubSub
	busSubLiveMatches chan interface{}
}

func NewUpdater(options *UpdaterOptions) *Updater {
	return &Updater{
		options: options,
		log:     options.Logger.WithPackage(processorName),
	}
}

func (p *Updater) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if p.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(p.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: shutdown,
	}
}

func (p *Updater) Start(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return nsproc.ErrProcessorContextDotaClient
	}

	if p.bus = nsctx.GetPubSub(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextPubSub
	}

	p.ctx = ctx
	p.busSubLiveMatches = p.bus.Sub(nsbus.TopicLiveMatches)
	p.activeReqs = make(map[uint32]bool)

	pool, err := ants.NewPool(p.options.PoolSize)

	if err != nil {
		p.log.WithError(err).Error("error starting worker pool")
		return err
	}

	p.workerPool = pool

	return p.loop()
}

func (p *Updater) profileCardExpired(c *models.PlayerProfileCard) bool {
	return p.db.NewRecord(c) || time.Since(c.UpdatedAt) > cardExpiration
}

func (p *Updater) loop() error {
	defer func() {
		p.workerPool.Release()
	}()

	p.log.Info("start")

	for {
		select {
		case busmsg, ok := <-p.busSubLiveMatches:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.(*nsbus.LiveMatchesDotaMessage); ok {
				p.handleLiveMatches(msg)
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *Updater) handleLiveMatches(matches *nsbus.LiveMatchesDotaMessage) {
	p.log.WithFields(logrus.Fields{
		"index": matches.Index,
		"count": len(matches.Matches),
	}).Debug("processing live matches")

	for _, match := range matches.Matches {
		if err := p.ctx.Err(); err != nil {
			p.log.
				WithField("match_id", match.GetMatchId()).
				WithError(err).
				Error("error processing live match")

			return
		}

		p.handlePlayers(match.GetPlayers())
	}
}

func (p *Updater) handlePlayers(players []*protocol.CSourceTVGameSmall_Player) {
	for _, livePlayer := range players {
		accountID := livePlayer.GetAccountId()
		plog := p.log.WithField("account_id", accountID)

		if err := p.ctx.Err(); err != nil {
			plog.WithError(err).Error("update profile card interrupted")
			return
		}

		err := p.workerPool.Submit(func() {
			if err := p.updateProfileCard(accountID); err != nil {
				plog.WithError(err).Error("error updating profile card")
			}
		})

		if err != nil {
			p.log.WithError(err).Error("error submiting task to worker pool")
		}
	}
}

func matchesProfileCardResponse(accountID uint32) func(proto.Message) bool {
	return func(resp proto.Message) bool {
		if body, ok := resp.(*protocol.CMsgDOTAProfileCard); ok {
			return body.GetAccountId() == accountID
		}

		return false
	}
}

func (p *Updater) requestProfileCard(accountID uint32) (*protocol.CMsgDOTAProfileCard, error) {
	req := &protocol.CMsgClientToGCGetProfileCard{
		AccountId: proto.Uint32(accountID),
	}

	resp := &protocol.CMsgDOTAProfileCard{}

	return resp, p.dota.MakeRequest(
		p.ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse),
		resp,
		matchesProfileCardResponse(accountID),
	)
}

func (p *Updater) updateProfileCard(accountID uint32) error {
	plog := p.log.WithField("account_id", accountID)

	if err := p.ctx.Err(); err != nil {
		plog.WithError(err).Error("update profile card interrupted")
		return err
	}

	card := &models.PlayerProfileCard{}
	result := p.db.Where(models.PlayerProfileCard{AccountID: accountID}).FirstOrInit(card)

	if result.Error != nil {
		plog.WithError(result.Error).Error("error creating profile card")
		return result.Error
	}

	if !p.profileCardExpired(card) {
		return nil
	}

	p.activeReqsMtx.Lock()
	requestPending := p.activeReqs[accountID]
	p.activeReqsMtx.Unlock()

	if requestPending {
		return nil
	}

	p.activeReqsMtx.Lock()
	p.activeReqs[accountID] = true
	p.activeReqsMtx.Unlock()

	resp, err := p.requestProfileCard(accountID)

	if err != nil {
		plog.WithError(err).Error("error requesting profile card")
		return err
	}

	if resp.GetAccountId() != accountID {
		p.log.WithError(ErrInvalidProfileCardResponse).WithFields(logrus.Fields{
			"expected": accountID,
			"received": resp.GetAccountId(),
		}).Error("received profile card for a different player")

		return ErrInvalidProfileCardResponse
	}

	update := models.PlayerProfileCardProto(resp)

	if p.db.NewRecord(card) {
		result = p.db.Create(update)
	} else {
		result = p.db.Model(models.PlayerProfileCardModel).Updates(update)
	}

	if result.Error != nil {
		plog.WithError(result.Error).Error("error updating profile card")
		return result.Error
	}

	p.activeReqsMtx.Lock()
	delete(p.activeReqs, accountID)
	p.activeReqsMtx.Unlock()

	return nil
}
