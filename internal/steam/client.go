package steam

import (
	"github.com/faceit/go-steam"
)

type Client struct {
	*steam.Client

	Session *Session
}

func NewClient(c *steam.Client) *Client {
	return &Client{
		Client:  c,
		Session: &Session{c: c},
	}
}
