package repository

import (
	"fmt"
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *repository) DeleteEmployeeRepository(employeeID string) error {
	employee := new(model.EmployeeDTO)
	if err := repo.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppError(http.StatusNotFound, "Employee not found", errors.ErrNotFound)
		}

		return err
	}

	if employee.ID == uuid.Nil {
		return errors.NewAppError(http.StatusNotFound, "Employee not found", errors.ErrNotFound)
	}

	if err := repo.db.Table(helpers.Employees).Where("id = ?", employeeID).Delete(&model.EmployeeDTO{}).Error; err != nil {
		return err
	}

	if err := repo.db.Table(helpers.Addresses).Where("employee_id = ?", employeeID).Delete(&model.AddressDTO{}).Error; err != nil {
		return err
	}

	if err := repo.DeleteFromCache(*employee); err != nil {
		fmt.Printf("Failed to invalidate cache: %v\n", err)
	}

	return nil
}
