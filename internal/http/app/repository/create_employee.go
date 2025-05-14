package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) CreateEmployeeRepository(employee model.EmployeeDTO) error {
	if err := repo.db.Table(helpers.Employees).Create(&employee).Error; err != nil {
		return err
	}

	err := repo.UpdateCache(employee)
	if err != nil {
		return err
	}

	return nil
}
