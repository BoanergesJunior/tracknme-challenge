package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) DeleteEmployeeRepository(employeeID string) error {
	return repo.db.Table(helpers.Employees).Where("id = ?", employeeID).Delete(&model.EmployeeDTO{}).Error
}
