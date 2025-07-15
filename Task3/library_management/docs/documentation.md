# Library Services Package

This is a consle based program to manage a **library system**, including:

- Adding and removing members
- Adding and removing books
- Borrowing and returning books
- Listing available and borrowed books

---

##  **Package Structure**

- **Library struct**
  - Manages all books and members.
  - Fields:
    - `Book`: map of books with their IDs as keys.
    - `Members`: map of members with their IDs as keys.

- **LibraryManager interface**
  - Defines the main operations for library management.

---

##  **Implemented Functions**

###  `Addmember(id int, name string) error`

Adds a new member with the given ID and name.

- **Parameters:**
  - `id`: unique integer ID for the member.
  - `name`: name of the member.
- **Returns:** error if the ID is already in use.

---

###  `Delmember(id int) error`

Deletes a member from the library.

- **Parameters:**
  - `id`: ID of the member to delete.
- **Returns:** error if the member does not exist or has borrowed books that are not returned.

---

###  `AddBook(book models.Book) error`

Adds a book to the library.

- **Parameters:**
  - `book`: `models.Book` struct containing book details.
- **Returns:** error (currently always nil; can be extended for duplicate checks).

---

###  `RemoveBook(bookID int) error`

Removes a book from the library if it is not borrowed.

- **Parameters:**
  - `bookID`: ID of the book to remove.
- **Returns:** error if book does not exist or is currently borrowed.

---

###  `BorrowBook(bookID int, membersID int) error`

Marks a book as borrowed by a member.

- **Parameters:**
  - `bookID`: ID of the book.
  - `membersID`: ID of the member borrowing the book.
- **Returns:** error if book is already borrowed or book does not exit or member does not exist.

---

###  `ReturnBook(bookID int, membersID int) error`

Marks a book as returned by a member.

- **Parameters:**
  - `bookID`: ID of the book.
  - `membersID`: ID of the member returning the book.
- **Returns:** error if the book was not borrowed or the member does not exist.

---

###  `ListAvailableBooks() []models.Book`

Returns a list of books that are **not currently borrowed**.

---

###  `ListBorrowedBooks() []models.Book`

Returns a list of books that are **currently borrowed**.

---
##  **Usage Example**
Welcome to the library
1, Add Member
2, Remove Member
3, Add Book
4, Remove Book
5, Borrow Book
6, Return Book
7, List Available Books
8, List Borrowed Books
9, Exit
enter your choice: 1
enter id of the member:
101
enter name of the member:
Abrham
successfully added the member

1, Add Member
2, Remove Member
...
enter your choice: 3
enter id of the book:
201
enter title of the book:
Go Programming
enter name of the author:
Alan A. A. Donovan
successfully added the book

1, Add Member
2, Remove Member
...
enter your choice: 5
enter id of the book:
201
enter id of the member:
101
successfully borrowed the book

1, Add Member
2, Remove Member
...
enter your choice: 9
