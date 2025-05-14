package usecase

import "github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"

func (uc *usecase) GetEmployee(employeeID string) (*model.EmployeeDTO, error) {
	employee, err := uc.repo.GetEmployeeRepository(employeeID)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
