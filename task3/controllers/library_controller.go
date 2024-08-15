package controllers

import (
    "fmt"
    "task3/models"
    "task3/services"
)

func RunLibraryConsole() {
    library := services.NewLibrary()
    
    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove a book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books")
        fmt.Println("7. Exit")
        fmt.Print("Choose an option: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var id int
            var title, author string
            fmt.Print("Enter Book ID: ")
            fmt.Scan(&id)
            fmt.Print("Enter Book Title: ")
            fmt.Scan(&title)
            fmt.Print("Enter Book Author: ")
            fmt.Scan(&author)

            book := models.Book{
                ID:     id,
                Title:  title,
                Author: author,
                Status: "Available",
            }
            library.AddBook(book)
            fmt.Println("Book added successfully!")

        case 2:
            var id int
            fmt.Print("Enter Book ID to remove: ")
            fmt.Scan(&id)
            library.RemoveBook(id)
            fmt.Println("Book removed successfully!")

        case 3:
            var bookID, memberID int
            fmt.Print("Enter Book ID to borrow: ")
            fmt.Scan(&bookID)
            fmt.Print("Enter Member ID: ")
            fmt.Scan(&memberID)

            err := library.BorrowBook(bookID, memberID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book borrowed successfully!")
            }

        case 4:
            var bookID, memberID int
            fmt.Print("Enter Book ID to return: ")
            fmt.Scan(&bookID)
            fmt.Print("Enter Member ID: ")
            fmt.Scan(&memberID)

            err := library.ReturnBook(bookID, memberID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book returned successfully!")
            }

        case 5:
            books := library.ListAvailableBooks()
            fmt.Println("Available Books:")
            for _, book := range books {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }

        case 6:
            var memberID int
            fmt.Print("Enter Member ID: ")
            fmt.Scan(&memberID)

            books := library.ListBorrowedBooks(memberID)
            fmt.Println("Borrowed Books:")
            for _, book := range books {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }

        case 7:
            fmt.Println("Exiting...")
            return

        default:
            fmt.Println("Invalid option. Please try again.")
        }
    }
}
