package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repo (repository) -> Data access / Database layer -> This is going to interact with the DB

type Repo struct {
	coll *mongo.Collection // this is going to represent 1 single MongoDB connection like a table.
}

func NewRepo (db *mongo.Database) *Repo {
	return &Repo{ // whenever we insert the first note, mongo is going to create a collection automatically
		coll : db.Collection("notes"),
	}
}

func (r *Repo) Create (ctx context.Context, note Note) (Note, error) { 
	// (r *Repo): This makes the function a METHOD, not a plain function.
	// Meaning in plain English: “This function belongs to Repo.”
	
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

func (r* Repo) List (ctx context.Context) ([] Note, error) { // returns list of note

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{} // bson.M is a map which is used to build our mongo filters (similar to how we write in nodejs: mongoose.find({ _condition_ }))
	// {} => means empty filter, means we want to return all docs.

	// .Find() -> returns a cursor (like an iterator) > over matching elements. It holds server resources.
	// Find executes a find command and returns a Cursor over the matching documents in the collection.
	cursor, err := r.coll.Find(childCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("Find notes failed: %w", err)
	} 

	// Cursor must be closed after use to avoid any kind of leaks
	defer cursor.Close(childCtx)

	var notes [] Note
	
	if err := cursor.All(childCtx, &notes); err != nil {
		return nil, fmt.Errorf("Decode notes failed: %w", err)
	}

	return notes, nil
}