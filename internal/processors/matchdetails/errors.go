package matchdetails

import (
	nserr "github.com/13k/night-stalker/internal/errors"
)

type errMatchSave struct {
	*nserr.Err
	MatchID uint64
}
