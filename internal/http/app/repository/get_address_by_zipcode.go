package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *repository) GetAddressByZipCode(employeeID uuid.UUID, zipCode string) (*model.AddressDTO, error) {
	addressCache, err := r.GetAddressByZipCodeCache(zipCode)
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	if addressCache != nil {
		return addressCache, nil
	}

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

func (r *repository) GetAddressByZipCodeCache(zipCode string) (*model.AddressDTO, error) {
	key := fmt.Sprintf("%s:%s", helpers.AddressKeyPrefix, zipCode)
	addressCache, err := r.GetZipCodeCache(context.Background(), key)
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	if addressCache != "" {
		var address model.AddressDTO
		if err := json.Unmarshal([]byte(addressCache), &address); err != nil {
			return nil, err
		}
		return &address, nil
	}

	return nil, errors.ErrNotFound
}
