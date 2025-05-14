package usecase

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *usecase) UpdateEmployeeFields(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	employeeDB, err := uc.repo.GetEmployeeRepository(employeeID)
	if err != nil {
		return nil, err
	}

	if employeeDB == nil {
		return nil, errors.NewAppError(http.StatusNotFound, "Employee not found", nil)
	}

	var addressUpdateTx *gorm.DB
	if employee.ZipCode != employeeDB.ZipCode {
		addressID, addressTx, err := uc.UpsertAddressDetails(uuid.MustParse(employeeID), &employee)
		if err != nil {
			return nil, err
		}
		employee.Address = addressID.String()
		addressUpdateTx = addressTx
	}

	updatedEmployee, err := uc.repo.UpdateEmployeeFieldsRepository(employeeID, employee)
	if err != nil {
		return nil, err
	}

	if addressUpdateTx != nil {
		if err := addressUpdateTx.Commit().Error; err != nil {
			return nil, err
		}
	}

	return updatedEmployee, nil
}
