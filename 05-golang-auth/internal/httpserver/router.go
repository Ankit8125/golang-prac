package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter (a *app.App) *gin.Engine {
	// Creating a new router instance
	r := gin.New() // New returns a new blank Engine instance without any middleware attached

	r.Use(gin.Logger()) // A Logger middleware that will write the logs to gin
	r.Use(gin.Recovery()) // A middleware that recovers from any panics (like server crash) and writes a 500 if there was one.

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)

	userSvc := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	// unauth route -> public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// List all data/files (protected)
	api := r.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))
	
	api.GET("/files", func(c *gin.Context) {
	
		userID, _ := middleware.GetUserID(c)

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"userid": userID,
			"files": [] any {},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"products": [] any {},
		})
	})

	admin := api.Group("admin")

	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func (c *gin.Context)  {
		role, _ := middleware.GetUserRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"role": role,
		})
	})

	return r
} 