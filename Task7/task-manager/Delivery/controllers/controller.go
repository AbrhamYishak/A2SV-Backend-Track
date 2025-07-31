package controllers

import (
	"net/http"
	"task7/Domain"
	"task7/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TaskController struct {
	TaskUseCase *Usecases.TaskUsecase
}

type TaskDTO struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	Duedate     time.Time          `json:"duedate" bson:"duedate"`
}

func NewTaskController(uuc *Usecases.TaskUsecase) *TaskController {
	return &TaskController{
		TaskUseCase: uuc,
	}
}
func (tc *TaskController) CreateTask(c *gin.Context){
	var task TaskDTO
	err := c.BindJSON(&task)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	if err = tc.TaskUseCase.CreateTask(tc.ChangeToTask(task)); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Succesfully created the task"})
}
func (tc *TaskController) GetTasks (c *gin.Context){
	tasks, err := tc.TaskUseCase.GetTasks()
	if err != nil{	
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}
func (tc *TaskController) GetByID (c *gin.Context){
	id := c.Param("id")
	task, err := tc.TaskUseCase.GetByID(id)
	if err != nil{	
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}
func(tc *TaskController) EditTaskByID(c *gin.Context){	
	id := c.Param("id")
    var t TaskDTO
    if err := c.BindJSON(&t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
	   return
    }
	err := tc.TaskUseCase.EditTask(id, tc.ChangeToTask(t))
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Succesfully edited the task"})	
}
func(tc *TaskController) DelTasksByID(c *gin.Context){	
	id := c.Param("id")
	err := tc.TaskUseCase.DeleteTask(id)
	if err != nil{ 	
	   c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
	   return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Successfully deleted the task"})
}
func (tc *TaskController) ChangeToTask (t TaskDTO) * Domain.Task{
    var task Domain.Task  
	task.ID = t.ID.Hex()
	task.Description = t.Description
	task.Title = t.Title
	task.Completed = t.Completed
	task.Duedate = t.Duedate
	return &task
}
