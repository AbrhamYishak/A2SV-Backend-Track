package routers

import (
	"log"
	"task7/Delivery/controllers"
	"github.com/gin-gonic/gin"
)
func StartRoute(uc *controllers.TaskController){
     r := gin.Default()
	 r.POST("/tasks", uc.CreateTask)
	 r.GET("/tasks",uc.GetTasks)
	 // r.POST("register", controllers.Register)
	 // r.POST("login", controllers.Login)
	 // r.Use(middleware.JwtAuthMiddlewareUser("your_jwt_secret"))
	 // r.GET("/tasks/:id",controllers.GetTasksByID)
	 // r.Use(middleware.JwtAuthMiddlewareAdmin("your_jwt_secret"))
	 // r.POST("/promote", controllers.Promote)
	 // r.DELETE("tasks/:id",controllers.DelTasksByID)
	 // r.PUT("tasks/:id", controllers.EditTasksByID)
	 if err := r.Run(); err != nil{
	    log.Fatal("could not start the server")	 
	 }
}
