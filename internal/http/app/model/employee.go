package model

import "github.com/google/uuid"

type EmployeeDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Age          int       `json:"age" validate:"required"`
	ZipCode      string    `json:"zip_code" validate:"required,len=8"`
	Gender       string    `json:"gender" validate:"oneof=M F"`
	Address      string    `json:"address"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state" validate:"len=2"`
}
