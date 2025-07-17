package controllers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task4/models"
	"task4/data"
)
func GetTasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, data.Store)
}
func GetTasksByID(c *gin.Context){
	id := c.Param("id")
	status,t:= data.GetByID(id); 
	if status==200{
	   c.IndentedJSON(status, t)
	   return
	}
	c.IndentedJSON(status, gin.H{"message":"Task not found with the given id"})

}
func CreateTasks(c *gin.Context){
   var t models.Task
   if err := c.BindJSON(&t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	   return
   }
   if err := data.ValidateForCreation(t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	   return
   } 
	status,err:= data.Addtask(t)
	if err!=nil{
	   c.IndentedJSON(status, gin.H{"message":err.Error()})
	   return
	}
	c.IndentedJSON(status, gin.H{"message":"Successfully created the task"})
}
func DelTasksByID(c *gin.Context){	
	id := c.Param("id")
	status, err := data.Deltask(id)
	if err != nil{ 	
	   c.IndentedJSON(status, gin.H{"message":err.Error()})
	   return
	}
	c.IndentedJSON(status, gin.H{"message":"Successfully deleted the task"})
}
func EditTasksByID(c *gin.Context){	
	id := c.Param("id")
    var t models.Task
    if err := c.BindJSON(&t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
	   return
    }
    if err := data.ValidateForEdit(t); err != nil{
	   c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	   return
    } 
	status , err := data.EditTask(t, id)
	if err != nil{
		c.IndentedJSON(status, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(status, gin.H{"message":"Successfully updated the task"})
}


