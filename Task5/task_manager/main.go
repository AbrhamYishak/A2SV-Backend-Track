package main
import (
	"task5/router"
	"task5/db"
)
func main(){
   db.Connection()	
   router.StartRoute() 
}
