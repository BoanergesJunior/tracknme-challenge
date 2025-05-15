package http

import (
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
	uc     model.IUsecase
}

func NewHandler(uc model.IUsecase) *Handler {
	router := gin.Default()

	router.Use(middleware.ErrorHandler())

	h := &Handler{
		router: router,
		uc:     uc,
	}

	h.Routes()
	return h
}

func (h *Handler) Routes() {
	h.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	employees := h.router.Group("/employees")
	{
		employees.POST("", h.CreateEmployee)
		employees.GET("", h.ListEmployees)
		employees.GET("/:id", h.GetEmployee)
		employees.GET("/zipcode/:zipCode", h.GetEmployeeByZipCode)
		employees.PUT("/:id", h.UpdateEmployee)
		employees.PATCH("/:id", h.UpdateEmployeeFields)
		employees.DELETE("/:id", h.DeleteEmployee)
	}
}

func (h *Handler) Start(addr string) error {
	return h.router.Run(addr)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
