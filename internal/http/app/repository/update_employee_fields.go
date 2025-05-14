package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (r *repository) UpdateEmployeeFieldsRepository(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	var employeeDB model.EmployeeDTO

	if err := r.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employeeDB).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]any)

	if employee.Name != "" {
		updates["name"] = employee.Name
	}
	if employee.Age != 0 {
		updates["age"] = employee.Age
	}
	if employee.ZipCode != "" {
		updates["zip_code"] = employee.ZipCode
	}
	if employee.Gender != "" {
		updates["gender"] = employee.Gender
	}
	if employee.Address != "" {
		updates["address"] = employee.Address
	}
	if employee.Neighborhood != "" {
		updates["neighborhood"] = employee.Neighborhood
	}
	if employee.City != "" {
		updates["city"] = employee.City
	}
	if employee.State != "" {
		updates["state"] = employee.State
	}

	if len(updates) == 0 {
		return &employeeDB, nil
	}

	if err := r.db.Table(helpers.Employees).Where("id = ?", employeeID).Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := r.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employeeDB).Error; err != nil {
		return nil, err
	}

	return &employeeDB, nil
}
