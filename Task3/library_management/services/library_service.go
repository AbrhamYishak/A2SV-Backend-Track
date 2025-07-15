package services

import (
	"task3/models"
	"errors"
	"fmt"
)

type Library struct{
    Book map[int]models.Book
    Members map[int]models.Member
}
type LibraryManager interface{
   Addmember(id int ,name string)error
   Delmember(id int)error
   AddBook(book models.Book)error
   RemoveBook(bookID int)
   BorrowBook(bookID int, membersID int)error 
   ReturnBook(bookID int, membersID int)error
   ListAvailableBooks()[]models.Book
   ListBorrowedBooks(membersID int)[]models.Book
}
func (l *Library) AddBook( book models.Book){
    l.Book[book.ID]	= book
}
func (l *Library) Addmember( id int, name string)error{
	_,ok := l.Members[id]
	if ok{
		return errors.New("the id already in use")
	}else{
		var m models.Member
		m.ID = id
		m.Name = name
		l.Members[id] = m
		return nil
	}
}
func (l *Library) Delmember( id int)error{
	b,ok := l.Members[id]
	if !ok{
		return errors.New("not member with this id")
	}else{
		if len(b.BorrowedBooks) != 0{
			return errors.New("could not remove a member without returning the books")
		}else{
			delete(l.Members, id)
			return nil
		}
	}
}
func (l *Library) RemoveBook(bookID int)error{
	b,ok := l.Book[bookID]
	if !ok{
		return errors.New("no book with the given id")
	}
	if b.Status == "borrowed"{
		return errors.New("could not remove borrowed it should be returned first")
	}
    delete(l.Book, bookID)
    return nil
}
func (l *Library) BorrowBook(bookID int, membersID int)error {
	b,ok := l.Book[bookID]
	if !ok{
		return errors.New("no book with the given id")
	}
    if b.Status == "not borrowed"{
        m,ok := l.Members[membersID] 
		if !ok{
			return errors.New("no member with this id")
		}
        b.Status = "borrowed"
        m.BorrowedBooks = append(m.BorrowedBooks, b) 
        l.Members[membersID] = m 
		l.Book[bookID] = b
		return nil
	}else{
		return errors.New("book already borrowed")
	}
}
func (l *Library) ReturnBook(bookID int, membersID int)error {
	b,ok := l.Book[bookID]
	if !ok{
		return errors.New("no book with the given id")
	}
    if b.Status == "borrowed"{
        m,ok := l.Members[membersID] 
		if !ok{
			return errors.New("no member with this id")
		}
        b.Status = "not borrowed"
        m.BorrowedBooks = append(m.BorrowedBooks, b) 
        l.Members[membersID] = m 
		l.Book[bookID] = b
		return nil
	}else{
		return errors.New("book not borrowed")
	}
}
func (l *Library) ListAvailableBooks( ) []models.Book{	
	var ans []models.Book
	for _,v := range l.Book{
		if v.Status == "not borrowed"{
			ans = append(ans,v)
		}
	}
	return ans
}

func (l *Library) ListBorrowedBooks( ) []models.Book{
	var ans []models.Book
	for _,v := range l.Book{
		if v.Status == "borrowed"{
			ans = append(ans,v)
		}
	}
	return ans
}
