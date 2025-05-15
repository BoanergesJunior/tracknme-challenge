package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *repository) UpsertAddressRepository(employeeID uuid.UUID, address model.AddressDTO) (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	existingAddress := new(model.AddressDTO)
	err := tx.Table(helpers.Addresses).Where("employee_id = ?", employeeID).First(&existingAddress).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			address.EmployeeID = employeeID
			if err := tx.Table(helpers.Addresses).Create(&address).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			tx.Rollback()
			return nil, err
		}
	} else {
		address.ID = existingAddress.ID
		address.EmployeeID = employeeID
		if err := tx.Table(helpers.Addresses).Save(&address).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	return tx, nil
}
