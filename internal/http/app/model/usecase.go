package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IEmployeeUsecase interface {
	CreateEmployee(employee EmployeeDTO) (EmployeeDTO, error)
	ListEmployees() ([]*EmployeeDTO, error)
	GetEmployee(employeeID string) (*EmployeeDTO, error)
	GetEmployeesByZipCode(zipCode string) ([]*EmployeeDTO, error)
	UpdateEmployee(employeeID string, employee EmployeeDTO) (*EmployeeDTO, error)
	UpdateEmployeeFields(employeeID string, employee EmployeeDTO) (*EmployeeDTO, error)
	DeleteEmployee(employeeID string) error
}

type IAddressUsecase interface {
	UpsertAddressDetails(employeeID uuid.UUID, employee *EmployeeDTO) (uuid.UUID, *gorm.DB, error)
}

type IUsecase interface {
	IEmployeeUsecase
	IAddressUsecase
}
