package tests

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) CreateEmployee(employee model.EmployeeDTO) (model.EmployeeDTO, error) {
	args := m.Called(employee)
	return args.Get(0).(model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) ListEmployees() ([]*model.EmployeeDTO, error)              { return nil, nil }
func (m *MockUsecase) GetEmployee(employeeID string) (*model.EmployeeDTO, error) { return nil, nil }
func (m *MockUsecase) GetEmployeesByZipCode(zipCode string) ([]*model.EmployeeDTO, error) {
	return nil, nil
}
func (m *MockUsecase) UpdateEmployee(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	return nil, nil
}
func (m *MockUsecase) UpdateEmployeeFields(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	return nil, nil
}
func (m *MockUsecase) DeleteEmployee(employeeID string) error { return nil }
func (m *MockUsecase) UpsertAddressDetails(employeeID uuid.UUID, employee *model.EmployeeDTO, address *model.AddressDTO) (*gorm.DB, error) {
	return nil, nil
}
