package middleware

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/gin-gonic/gin"
)

// ErrorHandler é um middleware que trata erros de forma consistente
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se há erros
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Verifica se é um AppError
			if appErr, ok := err.(*errors.AppError); ok {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
				return
			}

			// Erro genérico
			c.JSON(500, gin.H{
				"error": "Erro interno do servidor",
			})
		}
	}
}
