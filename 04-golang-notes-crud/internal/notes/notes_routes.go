package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes (r *gin.Engine, db *mongo.Database) {
	// create repo and handler once at startup
	repo := NewRepo(db)
	h := NewHandler(repo)

	notesGroup := r.Group("/notes") // This will be my parent node ("/notes")
	{
		notesGroup.POST("", h.CreateNote) // POST Req: -> /notes -> will run "CreateNote"
		notesGroup.GET("", h.ListNotes)
	}
}