package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *usecase) UpsertAddressDetails(employeeID uuid.UUID, employee *model.EmployeeDTO) (uuid.UUID, *gorm.DB, error) {
	formattedZipCode, err := formatZipCode(employee.ZipCode)
	if err != nil {
		return uuid.Nil, nil, err
	}

	existingAddress, err := uc.repo.GetAddressByZipCode(employeeID, formattedZipCode)
	if err != nil && err != errors.ErrNotFound {
		return uuid.Nil, nil, errors.NewAppError(
			http.StatusInternalServerError,
			"Erro ao verificar endereço existente",
			err,
		)
	}

	if existingAddress != nil && existingAddress.ZipCode == formattedZipCode {
		return existingAddress.ID, nil, nil
	}

	newAddress, err := uc.fetchAddressFromAPI(formattedZipCode)
	if err != nil {
		return uuid.Nil, nil, err
	}

	addressToSave := prepareAddressForUpsert(existingAddress, newAddress, employeeID, formattedZipCode)

	tx, err := uc.repo.UpsertAddressRepository(addressToSave)
	if err != nil {
		return uuid.Nil, nil, errors.NewAppError(
			http.StatusInternalServerError,
			"Erro ao salvar endereço",
			err,
		)
	}

	employee.City = newAddress.City
	employee.State = newAddress.State
	employee.Neighborhood = newAddress.Neighborhood

	return addressToSave.ID, tx, nil
}

func formatZipCode(zipCode string) (string, error) {
	cleanZipCode := strings.ReplaceAll(zipCode, "-", "")

	if len(cleanZipCode) != 8 {
		return "", errors.ErrInvalidZipCode
	}

	return fmt.Sprintf("%s-%s", cleanZipCode[:5], cleanZipCode[5:]), nil
}

func (uc *usecase) fetchAddressFromAPI(zipCode string) (*model.AddressDTO, error) {
	url := fmt.Sprintf("%s/ws/%s/json/", os.Getenv("ZIP_CODE_API"), zipCode)
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.NewAppError(
			http.StatusServiceUnavailable,
			"Erro ao consultar o serviço de CEP",
			err,
		)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, errors.ErrZipCodeNotFound
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.ErrZipCodeAPIError
	}

	var address model.AddressDTO
	if err := json.NewDecoder(response.Body).Decode(&address); err != nil {
		return nil, errors.NewAppError(
			http.StatusInternalServerError,
			"Erro ao processar resposta do serviço de CEP",
			err,
		)
	}

	if address.City == "" || address.State == "" {
		return nil, errors.ErrZipCodeNotFound
	}

	return &address, nil
}

func prepareAddressForUpsert(existingAddress *model.AddressDTO, newAddress *model.AddressDTO, employeeID uuid.UUID, zipCode string) model.AddressDTO {
	address := *newAddress
	address.EmployeeID = employeeID
	address.ZipCode = zipCode

	if existingAddress != nil {
		address.ID = existingAddress.ID
	} else {
		address.ID = uuid.New()
	}

	return address
}
