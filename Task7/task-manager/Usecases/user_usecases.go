package Usecases
import (
	"task7/Domain"
)
type UserUsecase struct{
	Repo UserRepoI
	PasswordService PasswordServiceI
}
func NewUserusecase (repo UserRepoI, Pas PasswordServiceI ) *UserUsecase{
	return &UserUsecase{
		Repo: repo,
		PasswordService: Pas,
	}
}
func (uu *UserUsecase) Register (*Domain.User) error{

}
