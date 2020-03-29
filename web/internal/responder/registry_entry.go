package responder

import (
	"sort"
)

type entry struct {
	r Responder
	p float32
}

type entries []*entry

func (s *entries) Add(r Responder, p float32) (int, *entry, bool) {
	t := r.MediaType().Type
	n := len(*s)
	i := sort.Search(n, func(i int) bool {
		return (*s)[i].p <= p
	})

	// updating existing
	if i < n {
		if e := (*s)[i]; e.r.MediaType().Type == t {
			e.p = p
			return i, e, false
		}
	}

	e := &entry{r: r, p: p}

	// appending new (including when empty)
	if i == n {
		*s = append(*s, e)
		return i, e, true
	}

	// prepending new
	if i == 0 {
		*s = append(entries{e}, *s...)
		return i, e, true
	}

	// inserting new
	*s = append(*s, nil)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = e

	return i, e, true
}

func (s entries) GetByType(t string) (int, *entry) {
	for i, e := range s {
		if e.r.MediaType().Type == t {
			return i, e
		}
	}

	return -1, nil
}
