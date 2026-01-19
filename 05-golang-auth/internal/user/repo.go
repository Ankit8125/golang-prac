package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repo handles all database operations for users (data access layer).
type Repo struct {
	col *mongo.Collection // Reference to the "users" collection in MongoDB
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		col: db.Collection("users"), // Returns existing collection or creates it if it doesn't exist
	}
}

// Create inserts a new user into the database and returns the user with the generated MongoDB ID.
func (r *Repo) Create(ctx context.Context, u User) (User, error) {
	// InsertOne adds the user document to MongoDB and returns a result with the generated _id
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return User{}, fmt.Errorf("Insert user failed: %w", err)
	}

	// MongoDB returns InsertedID as interface{} because it supports multiple ID types (string, int, ObjectID, etc).
	// Go requires explicit type conversion for type safety. We must assert it's an ObjectID before using it.
	// The comma-ok pattern safely performs the conversion: id = ObjectID value, ok = true if successful.
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("Insert user failed and inserted id is not object id: %w", err)
	} 

	u.ID = id
	return u, nil
}

func (r *Repo) FindByEmail (ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	
	filter := bson.M{"email": email}

	var u User

	err := r.col.FindOne(ctx, filter).Decode(&u); 
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments){ // If no data is present
			return User{}, mongo.ErrNoDocuments
		}

		return User{}, fmt.Errorf("Find by email failed: %w", err)
	}

	return u, nil
}