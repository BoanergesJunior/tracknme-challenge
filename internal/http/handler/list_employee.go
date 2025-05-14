package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListEmployees(c *gin.Context) {
	employees, err := h.uc.ListEmployees()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, employees)
}
