package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repo (repository) -> Data access layer -> This is where you are going to interact with the DB

type Repo struct {
	coll *mongo.Collection // this is going to represent 1 single MongoDB connection like a table.
}

func NewRepo (db *mongo.Database) *Repo {
	return &Repo{ // whenever we insert the first note, mongo is going to create a collection automatically
		coll : db.Collection("notes"),
	}
}

func (r *Repo) Create (ctx context.Context, note Note) (Note, error) {
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second) // context.WithTimeout() creates a new context that cancels after some time. Sometimes DB operations fails, so to avoid hanging of server, we are doing this.
	// ctx -> context from HTTP request
	// childCtx -> child context with additional 5 second timeout
	
	defer cancel()

	// Now we are going to insert our note in DB
	_, err := r.coll.InsertOne(childCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("Insert note failed")
	}

	return note, nil
}