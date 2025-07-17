package router

import (
	"task4/controllers"
	"github.com/gin-gonic/gin"
	"log"
)
func StartRoute(){
     r := gin.Default()
	 r.GET("/tasks",controllers.GetTasks)
	 r.GET("/tasks/:id",controllers.GetTasksByID)
	 r.POST("/tasks", controllers.CreateTasks)
	 r.DELETE("tasks/:id",controllers.DelTasksByID)
	 r.PUT("tasks/:id", controllers.EditTasksByID)
	 if err := r.Run(); err != nil{
	    log.Fatal("could not start the server")	 
	 }
}
