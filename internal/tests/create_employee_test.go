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

func TestCreateEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		payload        model.EmployeeDTO
		mockResponse   model.EmployeeDTO
		mockError      error
		expectedStatus int
		expectedBody   any
	}{
		{
			name: "successful employee creation",
			payload: model.EmployeeDTO{
				Name:    "Joao",
				Age:     30,
				ZipCode: "37026050",
				Gender:  "M",
				State:   "SP",
			},
			mockResponse: model.EmployeeDTO{
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
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   model.EmployeeDTO{},
		},
		{
			name: "invalid request body",
			payload: model.EmployeeDTO{
				Name:    "Joao",
				Age:     30,
				ZipCode: "123",
				Gender:  "X",
				State:   "ABC",
			},
			mockResponse:   model.EmployeeDTO{},
			mockError:      appErrors.ErrInvalidRequest,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"error": appErrors.ErrInvalidRequest.Message},
		},
		{
			name: "usecase error",
			payload: model.EmployeeDTO{
				Name:    "Joao",
				Age:     30,
				ZipCode: "37026050",
				Gender:  "M",
				State:   "SP",
			},
			mockResponse:   model.EmployeeDTO{},
			mockError:      appErrors.NewAppError(500, "internal server error", nil),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "internal server error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := new(MockUsecase)
			mockUC.On("CreateEmployee", tt.payload).Return(tt.mockResponse, tt.mockError)

			h := handler.NewHandler(mockUC)

			payloadBytes, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response model.EmployeeDTO
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.NotEmpty(t, response.ID)
				assert.Equal(t, tt.payload.Name, response.Name)
				assert.Equal(t, tt.payload.Age, response.Age)
				assert.Equal(t, tt.payload.ZipCode, response.ZipCode)
				assert.Equal(t, tt.payload.Gender, response.Gender)
				assert.Equal(t, tt.payload.State, response.State)
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
