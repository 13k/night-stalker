package responder

import (
	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

var registry = map[string]Responder{}

func init() {
	Register(JSON())
}

func Register(r Responder) {
	registry[r.MediaType().Type] = r
}

func GetByType(t string) Responder {
	return registry[t]
}

func GetByMediaType(t *nswebmime.MediaType) Responder {
	return GetByType(t.Type)
}

type Matcher interface {
	Match(string) bool
}

func GetByMatch(matcher Matcher) Responder {
	for t, r := range registry {
		if matcher.Match(t) {
			return r
		}
	}

	return nil
}
