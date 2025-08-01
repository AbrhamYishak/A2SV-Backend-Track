package Infrastructure
import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	"log"
	"fmt"
)
type Collection struct{
	task *mongo.Collection
	user *mongo.Collection
}
func initDB() * Collection{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}
	fmt.Println("Connected to MongoDB!")
    taskCollection := client.Database("taskdb").Collection("tasks")
    userCollection := client.Database("taskdb").Collection("users")
    return &Collection{
		task:  taskCollection,
		user:  userCollection,
	}
}
