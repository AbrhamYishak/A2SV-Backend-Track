package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
)
func display(name string ,store map[string]float64){
	 var ave float64
     fmt.Printf("+ %s  + %s +\n",strings.Repeat("-", 15),strings.Repeat("-", 5))
	 side := int((20-len(name))/2)
	 add := 0
	 if len(name)%2 == 1{
		 add = 1
	 }
	 fmt.Printf("| %s %s %s   |\n", strings.Repeat(" ", side+add), name, strings.Repeat(" ", side))
	 fmt.Printf("+ %s  + %s +\n",strings.Repeat("-", 15),strings.Repeat("-", 5))
	 fmt.Printf("| %-15s  | %-5s |\n", "Subject Name", "Score")
	 fmt.Printf("+ %s  + %s +\n",strings.Repeat("-", 15),strings.Repeat("-", 5))
     for k,v := range store{
		ave+=v
	    fmt.Printf("| %-15s  | %-5.2f |\n", k, v)
	    fmt.Printf("+ %s  + %s +\n",strings.Repeat("-", 15),strings.Repeat("-", 5))
	  }
     fmt.Printf("| %-15s  | %-5.2f |\n", "Average", Average(ave, len(store)))
	 fmt.Printf("+ %s  + %s +\n",strings.Repeat("-", 15),strings.Repeat("-", 5))
}
func Average(sum float64, l int) (ave float64){
	if l == 0{
		return 0
	}
	ave = sum/float64(l)
	return
}
func main(){
	var store = make(map[string]float64)
	var name string
	fmt.Println("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Println("Enter the number of subjects: ")
	var n int
	if _,err := fmt.Scan(&n); err!=nil{
		fmt.Println("you did not enter an integer")
		return
	}
	for i := range n{
		var name string
		var score float64
		fmt.Printf("Enter the %v subject name: ", i+1)
		if _,err := fmt.Scan(&name); err != nil{
			fmt.Println("you did not enter string")
		}
		fmt.Printf("Enter the %v subject score: ", i+1)
		if _,err := fmt.Scan(&score); err != nil{
			fmt.Println("you did not enter a number")
		}
		store[name] = score
	}
	display(name, store)
}
