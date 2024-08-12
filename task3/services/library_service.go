package services

import (
	"errors"
	"task3/models"
)

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library{
	return &Library{
		Books: make(map[int]models.Book),
		Members:make(map[int]models.Member),

	}

}

func (l *Library) AddBook(book models.Book){
	l.Books[book.ID]=book
}

func (l *Library) RemoveBook(bookID int){
	delete(l.Books,bookID)

}

func (l *Library) BorrowBook(bookID int,memberID int) error{
	book,exists:=l.Books[bookID]
	if !exists{
		return errors.New("book not found")

	}
	if book.Status !="Available"{
		return errors.New("book already borrowed")

	}
	member,exists:= l.Members[memberID]
	if !exists{
		return errors.New("member not found")
	}
	book.Status="Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks,book)
	l.Books[bookID]=book
	l.Members[memberID]=member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    member, exists := l.Members[memberID]
    if !exists {
        return errors.New("member not found")
    }

    for i, borrowedBook := range member.BorrowedBooks {
        if borrowedBook.ID == bookID {
            book := borrowedBook
            book.Status = "Available"
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            l.Books[bookID] = book
            l.Members[memberID] = member
            return nil
        }
    }

    return errors.New("book not found in borrowed books")
}

func (l *Library) ListAvailableBooks() []models.Book {
    var availableBooks []models.Book
    for _, book := range l.Books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, exists := l.Members[memberID]
    if !exists {
        return nil
    }
    return member.BorrowedBooks
}