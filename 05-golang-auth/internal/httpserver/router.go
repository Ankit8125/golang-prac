package httpserver

import "github.com/gin-gonic/gin"

func NewRouter () *gin.Engine {
	// Creating a new router instance
	r := gin.New() // New returns a new blank Engine instance without any middleware attached

	r.Use(gin.Logger()) // A Logger middleware that will write the logs to gin
	r.Use(gin.Recovery()) // A middleware that recovers from any panics (like server crash) and writes a 500 if there was one.

	r.GET("/health", health)

	return r
}