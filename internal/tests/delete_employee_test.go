package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	appErrors "github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	handler "github.com/BoanergesJunior/tracknme-challenge/internal/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		employeeID     string
		mockError      error
		expectedStatus int
		expectedBody   any
	}{
		{
			name:           "successful employee deletion",
			employeeID:     uuid.New().String(),
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   gin.H{"message": "Employee deleted successfully"},
		},
		{
			name:           "employee not found",
			employeeID:     uuid.New().String(),
			mockError:      appErrors.NewAppError(http.StatusNotFound, "Employee not found", nil),
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"error": "Employee not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := new(MockUsecase)
			mockUC.On("DeleteEmployee", tt.employeeID).Return(tt.mockError)

			h := handler.NewHandler(mockUC)

			req := httptest.NewRequest(http.MethodDelete, "/employees/"+tt.employeeID, nil)
			w := httptest.NewRecorder()

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if tt.expectedStatus == http.StatusOK {
				assert.Equal(t, tt.expectedBody.(gin.H)["message"], response["message"])
			} else {
				assert.Equal(t, tt.expectedBody.(gin.H)["error"], response["error"])
			}

			mockUC.AssertExpectations(t)
		})
	}
}
