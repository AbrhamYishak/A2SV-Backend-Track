package Usecases
import (
	"task7/Domain"
	"errors"
)
type UserUsecase struct{
	Repo UserRepoI
	PasswordService PasswordServiceI
	TokenService JwtI
}
func NewUserusecase (repo UserRepoI, Pas PasswordServiceI, ts JwtI ) *UserUsecase{
	return &UserUsecase{
		Repo: repo,
		PasswordService: Pas,
		TokenService: ts,
	}
}
func (uu *UserUsecase) Register (u *Domain.User) error{
	hashedpassword,err := uu.PasswordService.HashPassword(u.Password) 
	if err != nil{
		return err
	}
	_, exists := uu.Repo.UserExist(u.Username)
	if exists{
		return errors.New("the user already exists")
	}
	if u.Username == "" || u.Password == ""{
		return errors.New("username and password can not be empty")
	}
	count, err := uu.Repo.CountUsers() 
	if err != nil{
		return err
	}
	if count == 0{
		u.Isadmin = true
	}
	u.Password = hashedpassword
    err = uu.Repo.Register(u);
	return err
}
func (uu *UserUsecase) Login (lr *Domain.LoginRequest) (string,error){
	user, exists := uu.Repo.UserExist(lr.Username)
	if !exists{
		return "",errors.New("the user does not exists")
	}
	if lr.Username == "" || lr.Password == ""{
		return "",errors.New("the username and passoword can not be empty")
	}
	correct, err := uu.PasswordService.ComparePassword(lr.Password, user.Password)
	if err != nil|| !correct{
		return "",err
	}
	token, err := uu.TokenService.GenerateToken(&user)
	if err != nil{
		return "",err
	}
    return token,err
}

func (uu *UserUsecase) Promote (username string) error{
	return uu.Repo.Promote(username)
}
