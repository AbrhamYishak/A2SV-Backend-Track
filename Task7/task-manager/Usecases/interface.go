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
	Register(* Domain.User)error
	UserExist(username string)(Domain.User,bool)
	CountUsers()(int64, error)
	Isadmin(username string)(bool,error)
	// Promote(* Domain.User)error
}
type PasswordServiceI interface {
	HashPassword(password string) (string, error)
	ComparePassword(password string, hashedpassword string) (bool,error)
}
type JwtI interface{
	GenerateToken( user *Domain.User)(string,error)
	VerifyToken(token string)(bool,error)
	ExtractFromToken(token  string)(string, error)
}

