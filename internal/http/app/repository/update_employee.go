package repository

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
)

func (repo *repository) UpdateEmployeeRepository(employee model.EmployeeDTO, newAddress *model.AddressDTO) error {
	err := repo.updateEmployeeFields(employee.ID.String(), employee)
	if err != nil {
		return err
	}

	err = repo.updateAddress(employee.ID, newAddress)
	if err != nil {
		return err
	}

	err = repo.UpdateCache(employee, employee.ZipCode)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) updateEmployeeFields(employeeID string, employee model.EmployeeDTO) error {
	employeeDB := new(model.EmployeeDTO)
	if err := repo.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employeeDB).Error; err != nil {
		return err
	}

	if employeeDB.ID == uuid.Nil {
		return errors.NewAppError(http.StatusNotFound, "Employee not found", errors.ErrNotFound)
	}

	if err := repo.db.Table(helpers.Employees).Where("id = ?", employeeID).Updates(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (repo *repository) updateAddress(employeeID uuid.UUID, newAddress *model.AddressDTO) error {
	address := new(model.AddressDTO)
	if err := repo.db.Table(helpers.Addresses).Where("employee_id = ?", employeeID).First(&address).Error; err != nil {
		return err
	}

	if address.ID == uuid.Nil {
		return errors.NewAppError(http.StatusNotFound, "Address not found", errors.ErrNotFound)
	}

	if err := repo.db.Table(helpers.Addresses).Where("employee_id = ?", employeeID).Updates(&newAddress).Error; err != nil {
		return err
	}

	return nil
}
