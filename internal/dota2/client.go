package dota2

import (
	"github.com/paralin/go-dota2"
)

type Client struct {
	*dota2.Dota2

	Session *Session
}

func NewClient(d *dota2.Dota2) *Client {
	return &Client{
		Dota2:   d,
		Session: &Session{},
	}
}
