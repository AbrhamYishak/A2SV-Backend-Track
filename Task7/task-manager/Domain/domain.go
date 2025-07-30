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
	ID          string
	Username    string             
    Password    string             
	Isadmin     bool               
}
