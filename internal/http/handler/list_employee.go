package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListEmployees(c *gin.Context) {
	// TODO: Implementar a lógica de listagem usando o usecase
	c.JSON(http.StatusOK, gin.H{
		"message": "List employees endpoint",
	})
}
