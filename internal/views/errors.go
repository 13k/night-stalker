package views

import (
	"fmt"
)

type ErrAssociationNotEagerLoaded struct {
	Association string
}

func NewErrAssociationNotEagerLoaded(association string) *ErrAssociationNotEagerLoaded {
	return &ErrAssociationNotEagerLoaded{Association: association}
}

func (err *ErrAssociationNotEagerLoaded) Error() string {
	return fmt.Sprintf("association not eager-loaded: %s", err.Association)
}
