package controllers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task6/models"
	"task6/data"
    "golang.org/x/crypto/bcrypt"	
)
func GetTasks(c *gin.Context){
	status, t := data.Getdata()
	if status != 200{
		c.IndentedJSON(status, gin.H{"message":"could not find data"})
		return
	}
    c.IndentedJSON(status, t)	
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
	status , err := data.EditTask(id, t)
	if err != nil{
		c.IndentedJSON(status, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(status, gin.H{"message":"Successfully updated the task"})
}
func Register(c *gin.Context){
	var u models.User
	if err:=c.BindJSON(&u); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	if u.Username == "" || u.Password == ""{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"username and passowrd can not be empyt"})
		return
	}
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"could not hash password"})
		return
	}
	u.Password = string(HashedPassword)
	status, err := data.Register(u) 
	if err!=nil{
		c.IndentedJSON(status, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(status, gin.H{"message":"Succesfully Registered the user"})
}
func Login(c *gin.Context){
	var u models.User
	if err := c.BindJSON(&u); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	if u.Username == "" || u.Password == ""{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"username and passowrd can not be empyt"})
		return
	}
	status, token , err := data.Login(u)
	if err != nil{
		c.IndentedJSON(status, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(status, gin.H{"message":"Succesfully logged in","token":token})
}
func Promote(c *gin.Context){
	var u models.User
	if err:=c.BindJSON(&u); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid input"})
		return
	}
	if u.Username == ""{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"username can not be empyt"})
		return
	}
	status, err := data.Promote(u)
	if err != nil{
		c.IndentedJSON(status, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(status, gin.H{"message":"Succesfully promoted the user to Admin"})
}


