package main
import (
   "task7/Repositories"
   "task7/Usecases"
   "task7/Delivery/controllers"
   "task7/Delivery/routers"
   "task7/Infrastructure"
)
func main(){
    ps := Infrastructure.NewPasswordService()
	js := Infrastructure.NewJwtService()
	collection := Infrastructure.InitDB()
    trepo := Repositories.NewTaskRepo(collection.Task)
	urepo := Repositories.NewUserRepo(collection.User)
    tuc :=  Usecases.NewtaskUsecase(trepo)
	uuc := Usecases.NewUserusecase(urepo, ps, js)
    c := controllers.NewController(tuc, uuc)
    routers.StartRoute(c) 
}
