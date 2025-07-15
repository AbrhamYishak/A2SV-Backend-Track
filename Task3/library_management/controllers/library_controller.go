package controllers

import (
	"fmt"
	"task3/models"
	"task3/services"
)
func LibraryController(){
   l:= services.Library{}
   l.Members = make(map[int]models.Member)
   l.Book = make(map[int]models.Book)
   var choice int 
   fmt.Println("Welcome to the library") 
   for choice != 9{
   fmt.Println("1, Add Member")
   fmt.Println("2, Remove Member")
   fmt.Println("3, Add Book")
   fmt.Println("4, Remove Book")
   fmt.Println("5, Borrow Book")
   fmt.Println("6, Return Book")
   fmt.Println("7, List Available Books")
   fmt.Println("8, List Borrowed Books")
   fmt.Println("9, Exit")
   fmt.Println("enter your choice: ")
   if _,err := fmt.Scan(&choice); err!=nil{
	   fmt.Println("not valid input enter only integer")
   }
   switch choice {
   case 1: {
	   var name string
	   var id int
	   fmt.Println("enter id of the member: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   fmt.Println("enter name of the member: ")
       if _,err := fmt.Scan(&name); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   if err:=(l.Addmember(id,name)); err!=nil{
		   fmt.Println(err)
	   }
	   fmt.Println("successfully added the member")
   }
   case 2: {
	   var id int
	   fmt.Println("enter id of the member: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   if err:=(l.Delmember(id)); err!=nil{
		   fmt.Println(err)
	   }
	   fmt.Println("successfully Deleted the member")
   }
   case 3: {
	   var author string
	   var id int
	   var title string
	   status := "not borrowed" 
	   fmt.Println("enter id of the book: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   fmt.Println("enter title of the book: ")
       if _,err := fmt.Scan(&title); err!=nil{
	   fmt.Println("not valid input enter only strings")
       }
	   fmt.Println("enter name of the author: ")
       if _,err := fmt.Scan(&author); err!=nil{
	   fmt.Println("not valid input enter only strings")
	   }
	   b := models.Book{
		   ID: id,
		   Title: title,
		   Author: author,
		   Status: status,
	   }
	   l.AddBook(b)
	   fmt.Println("successfully added the book")
   }
   case 4: {
	   var id int
	   fmt.Println("enter id of the book: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   if err:=(l.RemoveBook(id)); err!=nil{
		   fmt.Println(err)
	   }
	   fmt.Println("successfully Removed the book")
   }
   case 5:{
	   var id int
	   fmt.Println("enter id of the book: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   var idM int
	   fmt.Println("enter id of the member: ")
       if _,err := fmt.Scan(&idM); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   if err:=(l.BorrowBook(id,  idM)); err!=nil{
		   fmt.Println(err)
	   }
	   fmt.Println("successfully borrowed the book")
   }
   case 6:{
	   var id int
	   fmt.Println("enter id of the book: ")
       if _,err := fmt.Scan(&id); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   var idM int
	   fmt.Println("enter id of the member: ")
       if _,err := fmt.Scan(&idM); err!=nil{
	   fmt.Println("not valid input enter only integer")
       }
	   if err:=(l.ReturnBook(id,  idM)); err!=nil{
		   fmt.Println(err)
	   }
	   fmt.Println("successfully returned the book")
   }
   case 7:{
       fmt.Println(l.ListAvailableBooks())
   }

   case 8:{
       fmt.Println(l.ListBorrowedBooks())
   }
   case 9:{
	   break
   }
   default:{
       fmt.Println("only enter from 1-9")
   }
}
}
}
