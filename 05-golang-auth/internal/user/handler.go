package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler (svc *Service) *Handler { // This is going to create a handler	 
	return &Handler{svc: svc} 
}

// Code flow: 
// Database Layer
// Business Logic (Calls to repo.go) -> repo.go talks to DB
// handler.go -> talks to service.go
// this handler.go will be used in route file

func (h *Handler) Register (c *gin.Context){
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	output, err := h.svc.Register(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *Handler) Login (c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	output, err := h.svc.Login(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}