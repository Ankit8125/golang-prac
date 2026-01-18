package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct { // This will be my model.
	// MongoDB stores data in BSON format
	ID        primitive.ObjectID `bson:"_id" json:"id"` // In bson, I am storing in "_id" format and in json I am storing in "id" format
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Pinned    bool               `bson:"pinned" json:"pinned"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required"` // 'binding' will tell 'gin' to validate this. So in case, it is empty, it's going to return an error.
	Content string `json:"content" binding:"required"`
	Pinned  bool   `json:"pinned"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Pinned  bool   `json:"pinned"`
}