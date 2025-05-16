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

func (m *MockUsecase) ListEmployees() ([]*model.EmployeeDTO, error) {
	args := m.Called()
	return args.Get(0).([]*model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) GetEmployee(employeeID string) (*model.EmployeeDTO, error) {
	args := m.Called(employeeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) GetEmployeesByZipCode(zipCode string) ([]*model.EmployeeDTO, error) {
	args := m.Called(zipCode)
	return args.Get(0).([]*model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) UpdateEmployee(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	args := m.Called(employeeID, employee)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) UpdateEmployeeFields(employeeID string, employee model.EmployeeDTO) (*model.EmployeeDTO, error) {
	args := m.Called(employeeID, employee)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.EmployeeDTO), args.Error(1)
}

func (m *MockUsecase) DeleteEmployee(employeeID string) error {
	args := m.Called(employeeID)
	return args.Error(0)
}

func (m *MockUsecase) UpsertAddressDetails(employeeID uuid.UUID, employee *model.EmployeeDTO, address *model.AddressDTO) (*gorm.DB, error) {
	args := m.Called(employeeID, employee, address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gorm.DB), args.Error(1)
}
