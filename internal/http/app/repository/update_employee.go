package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) UpdateEmployeeRepository(employeeID string, employee model.EmployeeDTO) error {
	repo.UpdateCache(employee)
	return repo.db.Table(helpers.Employees).Where("id = ?", employeeID).Updates(&employee).Error
}
