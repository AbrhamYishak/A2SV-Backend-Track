package router

import (
	"task6/controllers"	
	"task6/middleware"
	"github.com/gin-gonic/gin"
	"log"
)
func StartRoute(){
     r := gin.Default()
	 r.POST("register", controllers.Register)
	 r.POST("login", controllers.Login)
	 r.Use(middleware.JwtAuthMiddlewareUser("your_jwt_secret"))
	 r.GET("/tasks",controllers.GetTasks)
	 r.GET("/tasks/:id",controllers.GetTasksByID)
	 r.Use(middleware.JwtAuthMiddlewareAdmin("your_jwt_secret"))
	 r.POST("/promote", controllers.Promote)
	 r.POST("/tasks", controllers.CreateTasks)
	 r.DELETE("tasks/:id",controllers.DelTasksByID)
	 r.PUT("tasks/:id", controllers.EditTasksByID)
	 if err := r.Run(); err != nil{
	    log.Fatal("could not start the server")	 
	 }
}
