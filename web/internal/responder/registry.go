package responder

import (
	"sync"

	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

var defaultRegistry = &Registry{}

type Registry struct {
	entries entries
	mtx     sync.RWMutex
}

func (r *Registry) Register(res Responder, priority float32) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.entries.Add(res, priority)
}

func (r *Registry) GetByType(t string) Responder {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if _, e := r.entries.GetByType(t); e != nil {
		return e.r
	}

	return nil
}

func (r *Registry) GetByMediaType(t *nswebmime.MediaType) Responder {
	return r.GetByType(t.Type)
}

func (r *Registry) GetByMatch(m Matcher) Responder {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	for _, e := range r.entries {
		if m.Match(e.r.MediaType().Type) {
			return e.r
		}
	}

	return nil
}

func Register(r Responder, p float32) {
	defaultRegistry.Register(r, p)
}

func GetByType(t string) Responder {
	return defaultRegistry.GetByType(t)
}

func GetByMediaType(t *nswebmime.MediaType) Responder {
	return defaultRegistry.GetByMediaType(t)
}

func GetByMatch(m Matcher) Responder {
	return defaultRegistry.GetByMatch(m)
}
