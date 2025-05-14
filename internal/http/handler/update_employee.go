package http

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		c.Error(errors.ErrInvalidID)
		return
	}

	var employee *model.EmployeeDTO
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.Error(errors.NewAppError(
			errors.ErrInvalidRequest.Code,
			errors.ErrInvalidRequest.Message,
			err,
		))
		return
	}

	employee, err := h.uc.UpdateEmployee(employeeID, *employee)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, employee)
}
