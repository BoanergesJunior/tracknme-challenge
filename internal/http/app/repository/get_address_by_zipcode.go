package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *repository) GetAddressByZipCode(employeeID uuid.UUID, zipCode string) (*model.AddressDTO, error) {
	var address model.AddressDTO
	query := r.db.Table(helpers.Addresses).Where("employee_id = ?", employeeID)

	if zipCode != "" {
		query = query.Where("zip_code = ?", zipCode)
	}

	if err := query.First(&address).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return &address, nil
}
