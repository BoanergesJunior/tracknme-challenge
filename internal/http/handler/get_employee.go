package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get employee endpoint",
	})
}
