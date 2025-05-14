package usecase

import "github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"

func (uc *usecase) GetEmployeesByZipCode(zipCode string) ([]*model.EmployeeDTO, error) {
	employees, err := uc.repo.GetEmployeesByZipCodeRepository(zipCode)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
