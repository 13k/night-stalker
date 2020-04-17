package db

import (
	"fmt"

	nsm "github.com/13k/night-stalker/models"
)

type ErrMixedRecords struct {
	ExpectedModel nsm.Model
	InvalidModel  nsm.Model
}

func (err *ErrMixedRecords) Error() string {
	return fmt.Sprintf(
		"mixed models in records collection (expected %s, found %s)",
		err.ExpectedModel.Name(),
		err.InvalidModel.Name(),
	)
}
