package livematches

import (
	"math"

	"github.com/paralin/go-dota2/protocol"
)

type queryPage struct {
	index uint32
	start uint32
	total uint32
	res   *protocol.CMsgGCToClientFindTopSourceTVGamesResponse
}

type queryPageCacheKey [2]uint32

func newQueryPageCacheKey(page *queryPage) [2]uint32 {
	return [2]uint32{page.index, page.start}
}

// valid only for a single session.
// must be discarded/reset on session restart.
type queryPageCache map[queryPageCacheKey]*queryPage

func (c queryPageCache) Contains(page *queryPage) bool {
	return c[newQueryPageCacheKey(page)] != nil
}

func (c queryPageCache) Add(page *queryPage) {
	c[newQueryPageCacheKey(page)] = page
}

func (c queryPageCache) MaxIndex() uint32 {
	var index uint32

	for key := range c {
		if key[0] > index {
			index = key[0]
		}
	}

	return index
}

func (c queryPageCache) LatestPages() []*queryPage {
	maxIndex := c.MaxIndex()

	var pages []*queryPage

	for key, page := range c {
		if key[0] == maxIndex {
			pages = append(pages, page)
		}
	}

	return pages
}

type querySeq struct {
	index  uint32
	psize  uint32
	total  uint32
	npages int
	pages  queryPageCache
}

func newQuerySeq() *querySeq {
	return &querySeq{pages: make(queryPageCache)}
}

func (s *querySeq) Init(page *queryPage) {
	s.index = page.index
	s.psize = uint32(len(page.res.GetGameList()))
	s.total = page.res.GetNumGames()
	s.npages = int(math.Ceil(float64(s.total) / float64(s.psize)))
	s.pages = make(queryPageCache)
}

func (s *querySeq) Reset() {
	s.index = 0
	s.psize = 0
	s.total = 0
	s.npages = 0
	s.pages = make(queryPageCache)
}

func (s *querySeq) Cache(page *queryPage) {
	s.pages.Add(page)
}

func (s *querySeq) IsCached(page *queryPage) bool {
	return s.pages.Contains(page)
}

func (s *querySeq) IsEmpty() bool {
	return len(s.pages) == 0
}

func (s *querySeq) IsFull() bool {
	return len(s.Pages()) == s.npages
}

func (s *querySeq) MaxIndex() uint32 {
	return s.pages.MaxIndex()
}

func (s *querySeq) Pages() []*queryPage {
	return s.pages.LatestPages()
}
