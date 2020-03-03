package util

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/13k/night-stalker/models"
)

var (
	ErrFollowedPlayerAlreadyExists = errors.New("player already followed")
)

func FollowPlayer(db *gorm.DB, followed *models.FollowedPlayer, updateExisting bool) (*models.FollowedPlayer, error) {
	persisted := &models.FollowedPlayer{}
	scope := db.Where(&models.FollowedPlayer{AccountID: followed.AccountID})

	var result *gorm.DB

	if updateExisting {
		result = scope.Assign(followed).FirstOrCreate(persisted)
	} else {
		result = scope.Take(persisted)

		if err := result.Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				persisted = followed
				result = db.Create(persisted)
			}
		} else {
			return nil, ErrFollowedPlayerAlreadyExists
		}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return persisted, nil
}
