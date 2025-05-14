package usecase

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/google/uuid"
)

func (uc *usecase) CreateEmployee(employee model.EmployeeDTO) (model.EmployeeDTO, error) {
	employee.ID = uuid.New()

	addressID, addressTx, err := uc.UpsertAddressDetails(employee.ID, &employee)
	if err != nil {
		return model.EmployeeDTO{}, err
	}

	employee.Address = addressID.String()

	err = uc.repo.CreateEmployeeRepository(employee)
	if err != nil {
		return model.EmployeeDTO{}, err
	}

	if err := addressTx.Commit().Error; err != nil {
		return model.EmployeeDTO{}, err
	}

	return employee, nil
}
