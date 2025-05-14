package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) GetEmployeeRepository(employeeID string) (*model.EmployeeDTO, error) {
	employee := model.EmployeeDTO{}
	return &employee, repo.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employee).Error
}
