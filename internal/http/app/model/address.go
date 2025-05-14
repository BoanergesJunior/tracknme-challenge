package model

import "github.com/google/uuid"

type AddressDTO struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	EmployeeID   uuid.UUID `json:"employee_id" gorm:"index"`
	ZipCode      string    `json:"cep" gorm:"column:zip_code"`
	Street       string    `json:"logradouro"`
	Complement   string    `json:"complemento"`
	Unit         string    `json:"unidade"`
	Neighborhood string    `json:"bairro"`
	City         string    `json:"localidade"`
	State        string    `json:"uf"`
	StateName    string    `json:"estado" gorm:"column:state_name"`
	Region       string    `json:"regiao"`
	IBGECode     string    `json:"ibge" gorm:"column:ibge_code"`
	GIACode      string    `json:"gia" gorm:"column:gia_code"`
	AreaCode     string    `json:"ddd" gorm:"column:area_code"`
	SIAFICode    string    `json:"siafi" gorm:"column:siafi_code"`
}
