package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	handler "github.com/BoanergesJunior/tracknme-challenge/internal/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListEmployees(t *testing.T) {
	gin.SetMode(gin.TestMode)

	employees := []*model.EmployeeDTO{
		{
			ID:           uuid.New(),
			Name:         "Joao",
			Age:          30,
			ZipCode:      "37026050",
			Gender:       "M",
			State:        "SP",
			Address:      "Rua das Flores",
			Neighborhood: "Centro",
			City:         "Sao Paulo",
		},
		{
			ID:           uuid.New(),
			Name:         "Maria",
			Age:          25,
			ZipCode:      "37026051",
			Gender:       "F",
			State:        "RJ",
			Address:      "Rua das Palmeiras",
			Neighborhood: "Jardim",
			City:         "Rio de Janeiro",
		},
	}

	tests := []struct {
		name           string
		mockResponse   []*model.EmployeeDTO
		mockError      error
		expectedStatus int
		expectedBody   any
	}{
		{
			name:           "successful list employees",
			mockResponse:   employees,
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   []*model.EmployeeDTO{},
		},
		{
			name:           "empty list",
			mockResponse:   []*model.EmployeeDTO{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   []*model.EmployeeDTO{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := new(MockUsecase)
			h := handler.NewHandler(mockUC)

			mockUC.On("ListEmployees").Return(tt.mockResponse, tt.mockError).Once()

			req := httptest.NewRequest(http.MethodGet, "/employees", nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response []*model.EmployeeDTO
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if len(tt.mockResponse) > 0 {
				assert.Equal(t, len(tt.mockResponse), len(response))
				for i, employee := range tt.mockResponse {
					assert.Equal(t, employee.Name, response[i].Name)
					assert.Equal(t, employee.Age, response[i].Age)
					assert.Equal(t, employee.ZipCode, response[i].ZipCode)
					assert.Equal(t, employee.Gender, response[i].Gender)
					assert.Equal(t, employee.State, response[i].State)
					assert.Equal(t, employee.Address, response[i].Address)
					assert.Equal(t, employee.Neighborhood, response[i].Neighborhood)
					assert.Equal(t, employee.City, response[i].City)
				}
			} else {
				assert.Empty(t, response)
			}

			mockUC.AssertExpectations(t)
		})
	}
}
