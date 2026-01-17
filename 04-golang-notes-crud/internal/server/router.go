package server

import (
	"net/http"
	"notes-api/internal/notes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter (database *mongo.Database) *gin.Engine {
	r := gin.Default() // Creates a default instance with logger, recovery and middleware

	r.GET("/health", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H {
			"ok": true,
			"status": "healthy",
		})
	})

	notes.RegisterRoutes(r, database)

	return r
} 
