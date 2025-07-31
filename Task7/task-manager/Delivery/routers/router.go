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
	 r.GET("/tasks/:id",uc.GetByID)
	 r.PUT("tasks/:id", uc.EditTaskByID)
	 r.DELETE("tasks/:id",uc.DelTasksByID)
	 // r.POST("register", controllers.Register)
	 // r.POST("login", controllers.Login)
	 // r.Use(middleware.JwtAuthMiddlewareUser("your_jwt_secret"))
	 // r.Use(middleware.JwtAuthMiddlewareAdmin("your_jwt_secret"))
	 // r.POST("/promote", controllers.Promote)
	 if err := r.Run(); err != nil{
	    log.Fatal("could not start the server")	 
	 }
}
