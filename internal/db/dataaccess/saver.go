package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
)

type Saver struct {
	mq nsdb.ModelQueryer
}

func NewSaver(mqx nsdb.ModelQueryer) *Saver {
	return &Saver{mq: mqx}
}
