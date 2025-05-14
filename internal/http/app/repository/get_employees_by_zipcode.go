package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (r *repository) GetEmployeesByZipCodeRepository(zipCode string) ([]*model.EmployeeDTO, error) {
	employees := []*model.EmployeeDTO{}
	return employees, r.db.Table(helpers.Employees).Where("zip_code = ?", zipCode).Find(&employees).Error
}
