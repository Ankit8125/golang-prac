package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
In Node.js with Mongoose, you write: await mongoose.connect('mongodb://...');
And it handles everything internally. Go requires explicit resource management because Go doesn't have automatic garbage collection for non-memory resources (like database connections).
*/
 
// Refer Obsidian to get analogy of the below function.
func Connect (cfg config.Config) (*mongo.Client, *mongo.Database, error) {
	// This function returns three things: the client, the database, and an error.

	// Step 1: Create a Context with Timeout
	// This prevents your app from freezing if MongoDB is unreachable. After 10 seconds, MongoDB operations will timeout and fail gracefully.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Background returns a non-nil, empty Context.
	defer cancel()
	// context.Background() - Creates an empty, root context. It never cancels, has no values, and no deadline
	// context.WithTimeout() - Wraps that context with a 10-second deadline
	// defer cancel() - Critical! This cancels the context when the function exits (even on error)

	// Step 2: Connect to Client Options
	clientOptions := options.Client().ApplyURI(cfg.MONGOURI)
	// This configures the MongoDB connection settings using your connection URI (like mongodb://localhost:27017).

	// Step 3: Establish Connection
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("Mongo connect failed")
	}
	// mongo.Connect() does NOT immediately connect to MongoDB. It creates a client object and queues the connection attempt
	// If something goes wrong (invalid URI, network issue at client creation), it returns an error. We return nil values and an error if it fails.

	// Step 4: Verify Connection Works (Ping)
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Mongo ping failed ")
	}
	// This is the actual connection verification! Ping() sends a test command to MongoDB. This confirms the connection is actually working before we return
	// If MongoDB is down or unreachable, this will fail. Again, uses the 10-second timeout context

	// Step 5: Get Database Reference
	database := client.Database(cfg.MONGODB)
	// Database() is a lightweight operation—it doesn't connect to that specific database yet. It just creates a reference you'll use for queries. Both client and database are returned for use in your app

	return client, database, nil
}

func Disconnect (client *mongo.Client) error {

	// Create a context with 5-second timeout (shorter than connect—cleanup should be faster)
	// Disconnect gracefully - closes all open connections
	// Return any error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}