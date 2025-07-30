package main
import (
   "task7/Repositories"
   "task7/Usecases"
   "task7/Delivery/controllers"
   "task7/Delivery/routers"
)
func main(){
    repo := Repositories.NewTaskRepo()
    tuc :=  Usecases.NewtaskUsecase(repo)
    tc := controllers.NewTaskController(tuc)
    routers.StartRoute(tc) 
}
