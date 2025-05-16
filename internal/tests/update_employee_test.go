package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	appErrors "github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	handler "github.com/BoanergesJunior/tracknme-challenge/internal/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		employeeID     string
		payload        model.EmployeeDTO
		mockResponse   *model.EmployeeDTO
		mockError      error
		expectedStatus int
		expectedBody   any
	}{
		{
			name:       "successful employee update",
			employeeID: uuid.New().String(),
			payload: model.EmployeeDTO{
				Name:    "Joao Updated",
				Age:     31,
				ZipCode: "37026051",
				Gender:  "M",
				State:   "RJ",
			},
			mockResponse: &model.EmployeeDTO{
				ID:           uuid.MustParse(uuid.New().String()),
				Name:         "Joao Updated",
				Age:          31,
				ZipCode:      "37026051",
				Gender:       "M",
				State:        "RJ",
				Address:      "Rua das Flores Updated",
				Neighborhood: "Centro",
				City:         "Rio de Janeiro",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   &model.EmployeeDTO{},
		},
		{
			name:       "employee not found",
			employeeID: uuid.New().String(),
			payload: model.EmployeeDTO{
				Name:    "Joao",
				Age:     30,
				ZipCode: "37026050",
				Gender:  "M",
				State:   "SP",
			},
			mockResponse:   nil,
			mockError:      appErrors.NewAppError(http.StatusNotFound, "Employee not found", nil),
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"error": "Employee not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := new(MockUsecase)
			if tt.mockError == nil {
				mockUC.On("UpdateEmployee", tt.employeeID, tt.payload).Return(tt.mockResponse, nil)
			} else {
				mockUC.On("UpdateEmployee", tt.employeeID, tt.payload).Return(nil, tt.mockError)
			}

			h := handler.NewHandler(mockUC)

			payloadBytes, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPut, "/employees/"+tt.employeeID, bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var response model.EmployeeDTO
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.NotEmpty(t, response.ID)
				assert.Equal(t, tt.mockResponse.Name, response.Name)
				assert.Equal(t, tt.mockResponse.Age, response.Age)
				assert.Equal(t, tt.mockResponse.ZipCode, response.ZipCode)
				assert.Equal(t, tt.mockResponse.Gender, response.Gender)
				assert.Equal(t, tt.mockResponse.State, response.State)
				assert.Equal(t, tt.mockResponse.Address, response.Address)
				assert.Equal(t, tt.mockResponse.Neighborhood, response.Neighborhood)
				assert.Equal(t, tt.mockResponse.City, response.City)
			} else {
				var response map[string]any
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(gin.H)["error"], response["error"])
			}

			mockUC.AssertExpectations(t)
		})
	}
}
