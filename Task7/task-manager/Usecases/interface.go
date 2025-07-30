package UseCases
import (
	"go_demo/Domain"
)

type TaskRepoI interface {
	CreateTasks(* Domain.Task)error
	GetTasks()[]Domain.Task
	GetByID()[]Domain.Task,error
	EditTask(* Domain.Task)error
    DeleteTask(* Domain.Task)error
}
type UserRepoI interface{
	RegisterUser(* Domain.User)error
	LoginUser(* Domain.User)error
    Isadmin(* Domain.User)bool,error
	Promote(* Domain.User)error
}
type PasswordServiceI interface {
	HashPassword(password string) string
	ComparePassword(password string, hashedpassword string) bool
}
type JwtI interface{
	GenerateToken()
}