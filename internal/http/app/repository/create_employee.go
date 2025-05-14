package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) CreateEmployeeRepository(employee model.EmployeeDTO) error {
	return repo.db.Table(helpers.Employees).Create(&employee).Error
}
