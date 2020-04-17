package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
)

type Loader struct {
	mq nsdb.ModelQueryer
}

func NewLoader(mqx nsdb.ModelQueryer) *Loader {
	return &Loader{mq: mqx}
}
