package util

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/13k/night-stalker/models"
)

var (
	ErrFollowedPlayerAlreadyExists = errors.New("player already followed")
)

func FollowPlayer(db *gorm.DB, accountID uint32, label string, updateExisting bool) (*models.FollowedPlayer, error) {
	followed := &models.FollowedPlayer{AccountID: accountID}
	criteria := *followed
	result := db.Where(criteria).FirstOrInit(followed)

	if result.Error != nil {
		return nil, result.Error
	}

	if !db.NewRecord(followed) && !updateExisting {
		return nil, ErrFollowedPlayerAlreadyExists
	}

	if db.NewRecord(followed) {
		followed.Label = label
		result = db.Create(followed)
	} else {
		update := models.FollowedPlayer{Label: label}
		result = db.Model(models.FollowedPlayerModel).Updates(update)
	}

	return followed, result.Error
}
