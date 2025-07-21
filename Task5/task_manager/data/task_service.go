package data

import (
	"context"
	"errors"
	"net/http"
	"task5/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

func InitTaskController(collection *mongo.Collection) {
	taskCollection = collection
}
func Getdata()(int, []models.Task){
    var tasks []models.Task 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return http.StatusInternalServerError, tasks
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return http.StatusInternalServerError, tasks
		}
		tasks = append(tasks, task)
	}
	return http.StatusOK, tasks
}
func GetByID(id string)(int, models.Task){
    objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return http.StatusBadRequest, models.Task{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var task models.Task
	err = taskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
		 	return http.StatusNotFound, models.Task{}
		}
		return http.StatusInternalServerError, models.Task{}
	}
    return http.StatusOK, task
}
func Addtask(t models.Task)(int, error){
	t.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := taskCollection.InsertOne(ctx, t)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Failed to create task")
	}
    return http.StatusCreated, nil
}
func Deltask(id string) (int, error){
    objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
	    return	http.StatusBadRequest, errors.New("Invalid ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		return http.StatusInternalServerError, errors.New("Failed to delete task")
	}
    return http.StatusOK, nil
}
func EditTask(id string, t models.Task) (int ,error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return http.StatusBadRequest, errors.New("Invalid ID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"completed":   t.Completed,
			"duedate":   t.Duedate,
		},
	}

	result, err := taskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil || result.MatchedCount == 0 {
		return http.StatusInternalServerError, errors.New("could not update the task")
	}

	return http.StatusOK, nil
}
