package models
import (
	"time"
)

type Task struct{
	ID string
	Title string
	Description string
	Status bool
	Duedate time.Time
}
