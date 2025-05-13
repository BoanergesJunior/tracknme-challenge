package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteEmployee remove um employee
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	// TODO: Implementar a lógica de remoção usando o usecase
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee deleted successfully",
		"id":      id,
	})
}
