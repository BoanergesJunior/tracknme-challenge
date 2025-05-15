package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IEmployeeRepository interface {
	CreateEmployeeRepository(employee EmployeeDTO) error
	ListEmployeesRepository() ([]*EmployeeDTO, error)
	GetEmployeeRepository(employeeID string) (*EmployeeDTO, error)
	GetEmployeesByZipCodeRepository(zipCode string) ([]*EmployeeDTO, error)
	UpdateEmployeeRepository(employee EmployeeDTO, newAddress *AddressDTO) error
	UpdateEmployeeFieldsRepository(employeeID string, employee EmployeeDTO) (*EmployeeDTO, error)
	DeleteEmployeeRepository(employeeID string) error
}

type IAddressRepository interface {
	UpsertAddressRepository(employeeID uuid.UUID, address AddressDTO) (*gorm.DB, error)
	GetAddressByZipCode(employeeID uuid.UUID, zipCode string) (*AddressDTO, error)
}

type ICacheRepository interface {
	UpdateCache(employee EmployeeDTO, oldZipCode ...string) error
	DeleteFromCache(employee EmployeeDTO) error
}

type IRepository interface {
	IEmployeeRepository
	IAddressRepository
	ICacheRepository
}
