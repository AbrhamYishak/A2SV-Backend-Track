package Domain
import (
	"time"
)

type Task struct {
	ID          string
	Title       string            
	Description string             
	Completed   bool               
	Duedate     time.Time          
}
type User struct {
	ID          primitive.ObjectID 
	Username    string             
    Password    string             
	Isadmin     bool               
}
type TaskRepo interface{

}