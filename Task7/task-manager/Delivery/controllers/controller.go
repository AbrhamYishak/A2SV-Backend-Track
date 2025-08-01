package controllers

import (
	"net/http"
	"task7/Domain"
	"task7/Usecases"

	"github.com/gin-gonic/gin"
)
type Controller struct {
	TaskUseCase *Usecases.TaskUsecase
	UserUseCase *Usecases.UserUsecase
}
func NewTaskController(tuc *Usecases.TaskUsecase,uuc *Usecases.UserUsecase) *Controller {
	return &Controller{
		TaskUseCase: tuc,
		UserUseCase: uuc,
	}
}
func (con *Controller) CreateTask(c *gin.Context){
	var task Domain.Task
	err := c.BindJSON(&task)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	if err = con.TaskUseCase.CreateTask(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Succesfully created the task"})
}
func (con *Controller) GetTasks (c *gin.Context){
	tasks, err := con.TaskUseCase.GetTasks()
	if err != nil{	
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}
func (con *Controller) GetByID (c *gin.Context){
	id := c.Param("id")
	task, err := con.TaskUseCase.GetByID(id)
	if err != nil{	
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}
func(con *Controller) EditTaskByID(c *gin.Context){	
	id := c.Param("id")
    var t Domain.Task
    if err := c.BindJSON(&t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
	   return
    }
	err := con.TaskUseCase.EditTask(id, &t)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Succesfully edited the task"})	
}
func(con *Controller) DelTasksByID(c *gin.Context){	
	id := c.Param("id")
	err := con.TaskUseCase.DeleteTask(id)
	if err != nil{ 	
	   c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
	   return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Successfully deleted the task"})
}
func(con *Controller) Register (c *gin.Context){
   var user Domain.User
   if err := c.BindJSON(&user); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
	   return
   }
   if err := con.UserUseCase.Register(&user); err != nil{
	   c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
	   return
   }
   c.IndentedJSON(http.StatusCreated, gin.H{"message":"Succesfully created the user"}) 
}
func(con *Controller) Login (c *gin.Context){
	var request Domain.LoginRequest
	if err := c.BindJSON(&request); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	token ,err := con.UserUseCase.Login(&request)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token":token,"message":"Succesfully logged in"})
}
