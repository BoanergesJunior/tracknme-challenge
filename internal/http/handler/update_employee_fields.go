package http

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateEmployeeFields(c *gin.Context) {
	employeeID := c.Param("id")
	var employee model.EmployeeDTO
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEmployee, err := h.uc.UpdateEmployeeFields(employeeID, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}
