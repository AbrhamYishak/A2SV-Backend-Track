package main
import (
	"fmt"
	// "os"
   "strings"	
   "unicode"
)
func trimer(a string)string{
	var ans strings.Builder
	for _,v := range a{
		if !unicode.IsPunct(v)&&!unicode.IsSpace(v){
			ans.WriteRune(v)
		}

	}
	return ans.String()

}
func palindrome(word string)bool{
	a := strings.ToLower(trimer(word))
	i := 0
	j := len(a)-1
	for i < j{
		if a[i] == a[j]{
			i+=1
			j-=1
		}else{
			return false
		}
	}
	return true
   
}
func counter(word string)map[string]int{
   a := strings.ToLower(trimer(word))
   ans := make(map[string]int)
   for _,v := range a{
      ans[string(v)] = ans[string(v)]+1
   }
   return ans 
}
func main(){
	var a int
	fmt.Println("Welcome")
	fmt.Println("1, word frequency counter")
	fmt.Println("2, palindrome checker")
	if _,err := fmt.Scan(&a); err!=nil{
		fmt.Println("wrong input")
	}
    if a == 1{
		var input string
		fmt.Println("Enter the word: ")
	    if _,err := fmt.Scan(&input); err!=nil{
		    fmt.Println("wrong input ")
	    }
		fmt.Println(counter(input))
	}else if a == 2{
		var input string
		fmt.Println("Enter the word: ")
	    if _,err := fmt.Scan(&input); err!=nil{
		    fmt.Println("wrong input ")
	    }
		fmt.Println(palindrome(input))
	}else{
		fmt.Println("wrong input")
	}
}
