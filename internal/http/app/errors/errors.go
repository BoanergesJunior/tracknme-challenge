package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrInvalidZipCode = NewAppError(http.StatusBadRequest, "CEP inválido", nil)
	ErrInvalidGender  = NewAppError(http.StatusBadRequest, "Gênero inválido", nil)
	ErrInvalidAge     = NewAppError(http.StatusBadRequest, "Idade inválida", nil)
	ErrInvalidState   = NewAppError(http.StatusBadRequest, "Estado inválido", nil)
	ErrInvalidID      = NewAppError(http.StatusBadRequest, "ID inválido", nil)
	ErrInvalidRequest = NewAppError(http.StatusBadRequest, "Requisição inválida", nil)

	ErrZipCodeNotFound = NewAppError(http.StatusNotFound, "CEP não encontrado", nil)
	ErrZipCodeAPIError = NewAppError(http.StatusServiceUnavailable, "Erro ao consultar o serviço de CEP", nil)

	ErrDatabaseError = NewAppError(http.StatusInternalServerError, "Erro interno do servidor", nil)
	ErrCreateFailed  = NewAppError(http.StatusInternalServerError, "Erro ao criar registro", nil)
	ErrUpdateFailed  = NewAppError(http.StatusInternalServerError, "Erro ao atualizar registro", nil)
	ErrDeleteFailed  = NewAppError(http.StatusInternalServerError, "Erro ao deletar registro", nil)
	ErrNotFound      = NewAppError(http.StatusNotFound, "Registro não encontrado", nil)
	ErrListFailed    = NewAppError(http.StatusInternalServerError, "Erro ao listar registros", nil)
)
