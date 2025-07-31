package Repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task7/Domain"
	"task7/Usecases"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type TaskRepo struct {
	Collection *mongo.Collection
	Context    context.Context
}
type TaskDTO struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	Duedate     time.Time          `json:"duedate" bson:"duedate"`
}
func NewTaskRepo() Usecases.TaskRepoI {
	col := initDB()
	ctx := context.Background()

	return &TaskRepo{
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
func (t *TaskRepo) CreateTasks(task *Domain.Task) error {
	_, err := t.Collection.InsertOne(t.Context, task)
	return err
}

func (tr *TaskRepo) ChangeToTask (t TaskDTO) * Domain.Task{
    var task Domain.Task  
	task.ID = t.ID.Hex()
	task.Description = t.Description
	task.Title = t.Title
	task.Completed = t.Completed
	task.Duedate = t.Duedate
	return &task
}
func (t *TaskRepo) GetTasks() ([]Domain.Task,error){
    var tasks []Domain.Task 

	data, err := t.Collection.Find(t.Context, bson.M{})
	if err != nil {
		return []Domain.Task{}, errors.New("could not find the task")
	}
	for data.Next(t.Context) {
		var task TaskDTO
		if err := data.Decode(&task); err != nil {
			return []Domain.Task{}, errors.New("could not find the task")
		}
		tasks = append(tasks, *t.ChangeToTask(task))
	}
	return tasks, nil
}

func (t *TaskRepo) GetByID(id string) (Domain.Task,error){
    objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.Task{}, errors.New("could not convert id to objID")
	}
	var task Domain.Task
	err = t.Collection.FindOne(t.Context, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Domain.Task{}, errors.New("could not find the task with the given id")
		}
		return Domain.Task{}, errors.New("error finding the task")
	}
    return task, nil
}
func (tr *TaskRepo) EditTask(id string, t *Domain.Task) (error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid ID")
	}
	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"completed":   t.Completed,
			"duedate":   t.Duedate,
		},
	}
	result, err := tr.Collection.UpdateOne(tr.Context, bson.M{"_id": objID}, update)
	if err != nil || result.MatchedCount == 0 {
		return errors.New("could not update the task")
	}

	return nil
}
func (tc *TaskRepo) DeleteTask (id string) error{
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid ID")
	}
	result, err := tc.Collection.DeleteOne(tc.Context, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		return errors.New("Failed to delete task")
	}
    return nil
}
