package repository

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *repository) UpdateEmployeeFieldsRepository(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	employeeDB, err := r.getEmployeeByID(tx, employeeID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	updates := r.buildUpdatesMap(employee)
	if len(updates) == 0 {
		tx.Rollback()
		return employeeDB, nil
	}

	if err := r.applyUpdates(tx, employeeID, updates); err != nil {
		tx.Rollback()
		return nil, err
	}

	updatedEmployee, err := r.getEmployeeByID(tx, employeeID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := r.handleCacheUpdate(employeeDB, updatedEmployee); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return updatedEmployee, nil
}

func (r *repository) getEmployeeByID(tx *gorm.DB, employeeID string) (*model.EmployeeDTO, error) {
	var employee model.EmployeeDTO
	if err := tx.Table(helpers.Employees).Where("id = ?", employeeID).First(&employee).Error; err != nil {
		return nil, err
	}

	if employee.ID == uuid.Nil {
		return nil, errors.NewAppError(http.StatusNotFound, "Employee not found", errors.ErrNotFound)
	}

	return &employee, nil
}

func (r *repository) buildUpdatesMap(employee model.EmployeeDTO) map[string]any {
	updates := make(map[string]any)

	fields := map[string]any{
		"name":         employee.Name,
		"age":          employee.Age,
		"zip_code":     employee.ZipCode,
		"gender":       employee.Gender,
		"address":      employee.Address,
		"neighborhood": employee.Neighborhood,
		"city":         employee.City,
		"state":        employee.State,
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v != "" {
				updates[field] = v
			}
		case int:
			if v != 0 {
				updates[field] = v
			}
		}
	}

	return updates
}

func (r *repository) applyUpdates(tx *gorm.DB, employeeID string, updates map[string]any) error {
	return tx.Table(helpers.Employees).Where("id = ?", employeeID).Updates(updates).Error
}

func (r *repository) handleCacheUpdate(oldEmployee, newEmployee *model.EmployeeDTO) error {
	if oldEmployee.ZipCode != "" && oldEmployee.ZipCode != newEmployee.ZipCode {
		if err := r.DeleteFromCache(*newEmployee); err != nil {
			return err
		}
	}
	return r.UpdateCache(*newEmployee, oldEmployee.ZipCode)
}
