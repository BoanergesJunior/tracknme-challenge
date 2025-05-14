package usecase

import "github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"

func (uc *usecase) ListEmployees() ([]*model.EmployeeDTO, error) {
	employees, err := uc.repo.ListEmployeesRepository()
	if err != nil {
		return nil, err
	}

	return employees, nil
}
