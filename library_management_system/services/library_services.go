package services

import (
	"errors"
	"library_management_system/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exist := l.Books[bookID]

	if !exist {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	l.Books[bookID] = book
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {

	member, exist := l.Members[memberID]

	if !exist {
		return errors.New("member not found")
	}

	book, exist := l.Books[bookID]

	if !exist {
		return errors.New("book not found")
	}

	if book.Status == "Available" {
		return errors.New(("book is not Borrowed"))
	}

	newBorrowed := []models.Book{}

	for _, book := range member.BorrowedBooks {
		if book.ID != bookID {
			newBorrowed = append(newBorrowed, book)
		}
	}

	member.BorrowedBooks = newBorrowed

	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}

	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {

	member, exist := l.Members[memberID]

	if !exist {
		return []models.Book{}
	}

	return member.BorrowedBooks
}
