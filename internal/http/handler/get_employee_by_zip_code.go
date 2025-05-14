package http

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEmployeeByZipCode(c *gin.Context) {
	zipCode := c.Param("zipCode")
	if zipCode == "" {
		c.Error(errors.ErrInvalidZipCode)
		return
	}

	employees, err := h.uc.GetEmployeesByZipCode(zipCode)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, employees)
}
