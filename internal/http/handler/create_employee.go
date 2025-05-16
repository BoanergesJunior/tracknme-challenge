package http

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee model.EmployeeDTO
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.Error(errors.NewAppError(
			errors.ErrInvalidRequest.Code,
			errors.ErrInvalidRequest.Message,
			err,
		))
		return
	}

	createdEmployee, err := h.uc.CreateEmployee(employee)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, createdEmployee)
}
