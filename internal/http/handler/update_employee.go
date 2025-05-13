package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update employee endpoint",
	})
}
