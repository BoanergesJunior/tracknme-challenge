package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) ListEmployeesRepository() ([]*model.EmployeeDTO, error) {
	employees := []*model.EmployeeDTO{}
	return employees, repo.db.Table(helpers.Employees).Find(&employees).Error
}
