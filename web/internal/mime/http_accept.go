package mime

import (
	"path"
	"sort"
	"strconv"
	"strings"
)

type HTTPAcceptMediaType struct {
	*MediaType

	IsPattern bool
	Quality   float32
}

func NewHTTPAcceptMediaType(t *MediaType) (*HTTPAcceptMediaType, error) {
	amt := &HTTPAcceptMediaType{
		MediaType: t,
		IsPattern: strings.ContainsRune(t.Type, '*'),
	}

	if s := t.Params["q"]; s != "" {
		q, err := strconv.ParseFloat(s, 32)

		if err != nil {
			return nil, err
		}

		amt.Quality = float32(q)
	}

	return amt, nil
}

func (t *HTTPAcceptMediaType) Match(typ string) bool {
	ok, _ := path.Match(t.Type, typ)
	return ok
}

type HTTPAcceptMediaTypes []*HTTPAcceptMediaType

func ParseHTTPAcceptMediaTypes(s string) HTTPAcceptMediaTypes {
	types := strings.Split(s, ",")
	mediaTypes := make([]*HTTPAcceptMediaType, 0, len(types))

	for _, t := range types {
		if mt, err := ParseMediaType(t); err == nil {
			if acceptMT, err := NewHTTPAcceptMediaType(mt); err == nil {
				mediaTypes = append(mediaTypes, acceptMT)
			}
		}
	}

	return mediaTypes
}

func (s HTTPAcceptMediaTypes) SortByQuality() {
	byQuality := &httpAcceptMediaTypesByQuality{s: s}
	sort.Sort(byQuality)
}

type httpAcceptMediaTypesByQuality struct {
	s HTTPAcceptMediaTypes
}

func (s *httpAcceptMediaTypesByQuality) Len() int { return len(s.s) }

func (s *httpAcceptMediaTypesByQuality) Less(i, j int) bool {
	return s.s[i].Quality > s.s[j].Quality
}

func (s *httpAcceptMediaTypesByQuality) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
