package Infrastructure
import (
    "golang.org/x/crypto/bcrypt"	
	"errors"
	"task7/Usecases"
)
type PasswordService struct{

}
func NewPasswordService() Usecases.PasswordServiceI{
	return &PasswordService{}
}
func (ps *PasswordService) HashPassword (password string) (string,error){
    HashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err !=nil{
		return "", errors.New("could not encrypt the password")
	}
	return string(HashedPassword), nil
}
