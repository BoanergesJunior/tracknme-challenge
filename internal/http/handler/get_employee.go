package http

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		c.Error(errors.ErrInvalidID)
		return
	}

	employee, err := h.uc.GetEmployee(employeeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, employee)
}
