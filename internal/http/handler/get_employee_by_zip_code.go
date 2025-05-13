package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEmployeeByZipCode(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get employee by zipcode endpoint",
	})
}
