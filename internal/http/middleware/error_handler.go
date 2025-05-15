package middleware

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*errors.AppError); ok {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
				return
			}

			c.JSON(500, gin.H{
				"error": "Erro interno do servidor",
			})
		}
	}
}
