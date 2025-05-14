package http

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		c.Error(errors.ErrInvalidID)
		return
	}

	err := h.uc.DeleteEmployee(employeeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
