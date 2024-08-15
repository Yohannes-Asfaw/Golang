Library Management System Documentation
Overview
The Library Management System is a console-based application developed in Go that allows users to manage a collection of books and members of a library. It provides functionality to add, remove, borrow, and return books, as well as list available books and borrowed books.

Folder Structure
go
Copy code
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
main.go
Description: Entry point of the application. It initializes the application and runs the console interface.
Usage: go run main.go
controllers/library_controller.go
Description: Contains functions to handle console input and interact with the Library service methods.
Key Functions:
RunLibraryConsole(): Manages the console interface for user interaction.
models/book.go
Description: Defines the Book struct used in the application.
Fields:
ID: Identifier for the book (int)
Title: Title of the book (string)
Author: Author of the book (string)
Status: Status of the book ("Available" or "Borrowed") (string)
models/member.go
Description: Defines the Member struct used in the application.
Fields:
ID: Identifier for the member (int)
Name: Name of the member (string)
BorrowedBooks: Slice of books borrowed by the member ([]Book)
services/library_service.go
Description: Implements the LibraryManager interface and contains the business logic for managing books and members.
Key Types:
Library: The main struct that implements the LibraryManager interface.
Key Methods:
NewLibrary() *Library: Creates and returns a new Library instance.
AddBook(book Book): Adds a new book to the library.
RemoveBook(bookID int): Removes a book from the library by its ID.
BorrowBook(bookID int, memberID int) error: Allows a member to borrow a book if available.
ReturnBook(bookID int, memberID int) error: Allows a member to return a borrowed book.
ListAvailableBooks() []Book: Lists all available books in the library.
ListBorrowedBooks(memberID int) []Book: Lists all books borrowed by a specific member.
How to Run
Build and Run:

To build and run the application, use the following commands:
bash
Copy code
go mod tidy
go run main.go
Interact with Console:

Follow the prompts in the console to add, remove, borrow, return books, or list books.
Usage
Adding a Book
Option: 1
Input:
Book ID
Book Title
Book Author
Removing a Book
Option: 2
Input:
Book ID to remove
Borrowing a Book
Option: 3
Input:
Book ID to borrow
Member ID
Returning a Book
Option: 4
Input:
Book ID to return
Member ID
Listing Available Books
Option: 5
Listing Borrowed Books
Option: 6
Input:
Member ID
Exiting the Application
Option: 7
Error Handling
The application provides error messages for scenarios such as:
Book not found
Book already borrowed
Member not found
Book not found in borrowed books
Dependencies
Go 1.18 or higher
Future Enhancements
Add persistent storage (e.g., database integration)
Implement user authentication and authorization
Expand features to support reservations and fines