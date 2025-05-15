package usecase

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/google/uuid"
)

func (uc *usecase) CreateEmployee(employee model.EmployeeDTO) (model.EmployeeDTO, error) {
	employee.ID = uuid.New()

	addressTx, err := uc.UpsertAddressDetails(employee.ID, &employee, nil)
	if err != nil {
		return model.EmployeeDTO{}, err
	}

	err = uc.repo.CreateEmployeeRepository(employee)
	if err != nil {
		return model.EmployeeDTO{}, err
	}

	if addressTx != nil {
		if err := addressTx.Commit().Error; err != nil {
			return model.EmployeeDTO{}, err
		}
	}

	return employee, nil
}
