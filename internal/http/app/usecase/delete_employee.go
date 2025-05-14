package usecase

func (uc *usecase) DeleteEmployee(employeeID string) error {
	err := uc.repo.DeleteEmployeeRepository(employeeID)
	if err != nil {
		return err
	}

	return nil
}
