package Usecases
import (
	"task7/Domain"
)

type TaskRepoI interface {
	CreateTasks(* Domain.Task)error
	GetTasks()([]Domain.Task,error)
	GetByID(id string)(Domain.Task,error)
	EditTask(id string, t * Domain.Task)error
    DeleteTask(id string)error
}
type UserRepoI interface{
	RegisterUser(* Domain.User)error
	// LoginUser(* Domain.User)error
	//    Isadmin(* Domain.User)(bool,error)
	// Promote(* Domain.User)error
}
type PasswordServiceI interface {
	HashPassword(password string) (string, error)
	// ComparePassword(password string, hashedpassword string) bool
}
type JwtI interface{
	GenerateToken( user *Domain.User, secret string)(string)
	VerifyToken(token string)(bool,error)
    ExtractFromToken(token  string)(string, error)
}

