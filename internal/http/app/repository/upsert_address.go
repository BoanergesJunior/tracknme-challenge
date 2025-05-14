package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"gorm.io/gorm"
)

func (r *repository) UpsertAddressRepository(address model.AddressDTO) (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Table(helpers.Addresses).Save(&address).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return tx, nil
}
