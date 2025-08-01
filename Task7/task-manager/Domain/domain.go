package Domain
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	Duedate     time.Time          `json:"duedate" bson:"duedate"`
}
type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string  `json:"username" bson:"username"`           
	Password    string  `json:"password" bson:"password"`           
	Isadmin     bool    `json:"isadmin" bson:"isadmin"` 
}
type LoginRequest struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
