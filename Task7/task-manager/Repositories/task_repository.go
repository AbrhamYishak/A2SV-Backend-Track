package db
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
	"fmt"
	"time"
)
type TaskRepo struct {
	Collection *mongo.Collection
	Context    context.Context
}
func NewUserRepo() UseCase.IUserRepo {
	col := initDB(username)
	ctx := context.Background()

	return &UserRepo{
		Collection: col,
		Context:    ctx,
	}
}

func initDB() *mongo.Collection{
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
    return taskCollection
}