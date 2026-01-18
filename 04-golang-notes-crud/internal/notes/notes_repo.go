package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo (repository) -> Data access / Database layer -> This is going to interact with the DB

type Repo struct {
	coll *mongo.Collection // this is going to represent 1 single MongoDB connection like a table.
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{ // whenever we insert the first note, mongo is going to create a collection automatically
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {
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

func (r *Repo) List(ctx context.Context) ([]Note, error) { // returns list of note

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

	var notes []Note

	if err := cursor.All(childCtx, &notes); err != nil {
		return nil, fmt.Errorf("Decode notes failed: %w", err)
	}

	return notes, nil
}

func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id} // in bson we are storing as "_id", but in postman we are seeing as "id"

	var note Note

	err := r.coll.FindOne(childCtx, filter, options.FindOne()).Decode(&note) // FindOne doesn't return a cursor. It returns a single data
	if err != nil {
		return Note{}, fmt.Errorf("Find note by id failed: %w", err)
	}

	return note, nil
}

func (r *Repo) UpdateByID(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (Note, error) { // Now we also have fields while updating, so that we will get by "req" and it's type is UpdateNoteRequest
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id} // in bson we are storing as "_id", but in postman we are seeing as "id"

	update := bson.M{
		"$set": bson.M{
			"title":     req.Title,
			"content":   req.Content,
			"pinned":    req.Pinned,
			"UpdatedAt": time.Now().UTC(),
		},
	}

	after := options.After // hover on "after" -> you will see that it will return "options.ReturnDocument"

	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var updated Note

	err := r.coll.FindOneAndUpdate(childCtx, filter, update, &opts).Decode(&updated)
	if err != nil {
		return Note{}, fmt.Errorf("Update note failed: %w", err)
	}

	return updated, nil
}
