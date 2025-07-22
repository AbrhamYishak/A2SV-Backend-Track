package main
import (
	"task6/router"
	"task6/db"
)
func main(){
   db.Connection()	
   router.StartRoute() 
}
