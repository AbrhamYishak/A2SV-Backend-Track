package main
import (
   "task7/Repositories"
   "task7/Usecases"
   "task7/Delivery/controllers"
   "task7/Delivery/routers"
)
func main(){
    trepo := Repositories.NewTaskRepo()
	urepo := Repositories.NewUserRepo()
    tuc :=  Usecases.NewtaskUsecase(trepo)
	uuc := Usecases.NewUserUsecase(urepo)
    tc := controllers.NewTaskController(tuc)
	uc := controllers.NewUserController(uuc)
    routers.StartRoute(tc,uc) 
}
