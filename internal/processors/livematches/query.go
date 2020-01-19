package livematches

import (
	"sync"

	"github.com/paralin/go-dota2/protocol"
)

type queryPage struct {
	index uint32
	start uint32
	total uint32
	psize int
	res   *protocol.CMsgGCToClientFindTopSourceTVGamesResponse
}

func newQueryPage(res *protocol.CMsgGCToClientFindTopSourceTVGamesResponse) *queryPage {
	return &queryPage{
		index: res.GetGameListIndex(),
		start: res.GetStartGame(),
		total: res.GetNumGames(),
		psize: len(res.GetGameList()),
		res:   res,
	}
}

type discoveryPage struct {
	*queryPage

	mtx           sync.RWMutex
	lastPageStart uint32
}

func (p *discoveryPage) Empty() bool {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.queryPage == nil
}

func (p *discoveryPage) SetPage(page *queryPage) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.queryPage = page
	p.lastPageStart = p.total - uint32(p.psize)
}

func (p *discoveryPage) LastPageStart() uint32 {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.lastPageStart
}
