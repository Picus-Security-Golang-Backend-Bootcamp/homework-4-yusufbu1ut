package helpers

import (
	"errors"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/repositories"
)

//BookItem and AuthorItem is using on GETing and POSTing
type BookItem struct {
	models.Book
	Auth []models.Author `json:"authors"`
}
type AuthorItem struct {
	models.Author
	Book []models.Book `json:"books"`
}

func ListBook(bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem
	books, err := bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	for _, b := range books {
		var item BookItem
		item.Book = b
		book_authors, err := bookAuthorRepository.FindByISBN(b.ISBN)
		if err != nil {
			return nil, err
		}
		for _, ba := range book_authors {
			author, err := authorRepository.FindByID(int(ba.AuthorID))
			if err != nil {
				return nil, err
			}
			item.Auth = append(item.Auth, author[0])
		}
		items = append(items, item)
	}
	return items, nil
}

func ListAuth(bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem
	authors, err := authorRepository.GetAllAuthors()
	if err != nil {
		return nil, err
	}
	for _, a := range authors {
		var item AuthorItem
		item.Author = a
		book_authors, err := bookAuthorRepository.FindByAuthorID(int(a.ID))
		if err != nil {
			return nil, err
		}
		for _, ba := range book_authors {
			book, err := bookRepository.FindByID(ba.BookID)
			if err != nil {
				return nil, err
			}
			item.Book = append(item.Book, book[0])
		}
		items = append(items, item)
	}
	return items, nil
}

// searchByInput takes input parameter and first checks books and books' authors if it cant find any searchs on authors
func SearchByBookInput(srch string, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem
	books, err := bookRepository.FindByName(srch)
	if err != nil {
		return nil, err
	}
	for _, b := range books {
		var item BookItem
		item.Book = b
		book_authors, err := bookAuthorRepository.FindByISBN(b.ISBN)
		if err != nil {
			return nil, err
		}
		for _, ba := range book_authors {
			author, err := authorRepository.FindByID(int(ba.AuthorID))
			if err != nil {
				return nil, err
			}
			item.Auth = append(item.Auth, author[0])
		}
		items = append(items, item)
	}
	return items, nil
}

func SearchByAuthorInput(srch string, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem
	authors, err := authorRepository.FindByAuthorName(srch)
	if err != nil {
		return nil, err
	}
	for _, a := range authors {
		var item AuthorItem
		item.Author = a
		book_authors, err := bookAuthorRepository.FindByAuthorID(int(a.ID))
		if err != nil {
			return nil, err
		}
		for _, ba := range book_authors {
			book, err := bookRepository.FindByISBN(ba.BookID)
			if err != nil {
				return nil, err
			}
			item.Book = append(item.Book, book[0])
		}
		items = append(items, item)
	}
	return items, nil
}

//buyWithID works on books and takes id removes on its amount and saves it
func BuyWithID(id int, cnt int, bookRepository repositories.BookRepository) error {
	book, err := bookRepository.FindByID(id)
	if err != nil {
		return err
	}
	if len(book) == 0 {
		return errors.New("No Data with given id")
	}
	err = bookRepository.Buy(book[0], cnt)
	if err != nil {
		return err
	}
	return nil
}

// deleteByID fun takes int count and deletes which as connected bases
func DeleteByBookID(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	book, err := bookRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	_, err = bookAuthorRepository.DeleteByISBN(book[0].ISBN)
	if err != nil {
		return err
	}
	// //The authors that have no data in books are deleting too if he/she has no book, if delete this part only books will be deleted
	// for _, a := range authors {
	// 	err = authorRepository.DeleteByID(a)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func DeleteByAuthorID(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	err := authorRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	_, err = bookAuthorRepository.DeleteByAuthorID(id)
	if err != nil {
		return err
	}
	return nil
}

func GetBookById(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem = make([]BookItem, 1)
	var item BookItem
	book, err := bookRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	item.Book = book[0]
	ba, err := bookAuthorRepository.FindByISBN(book[0].ISBN)
	if err != nil {
		return nil, err
	}
	for _, a := range ba {
		author, err := authorRepository.FindByID(int(a.AuthorID))
		if err != nil {
			return nil, err
		}
		item.Auth = append(item.Auth, author[0])
	}
	items = append(items, item)
	return items, nil
}

func GetAuthById(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem = make([]AuthorItem, 1)
	var item AuthorItem
	auth, err := authorRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	item.Author = auth[0]
	ba, err := bookAuthorRepository.FindByAuthorID(id)
	if err != nil {
		return nil, err
	}
	for _, b := range ba {
		book, err := bookRepository.FindByISBN(b.BookID)
		if err != nil {
			return nil, err
		}
		item.Book = append(item.Book, book[0])
	}
	items = append(items, item)
	return items, nil
}

func CreateBook(b BookItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	// If we want, we can add other logical errors in if
	if b.ISBN <= 0 || b.Name == "" || b.StockAmount < 0 {
		return errors.New("Requested Fields name and isbn..")
	}
	err := bookRepository.Create(b.Book)
	if err != nil {
		return err
	}
	for _, ca := range b.Auth {
		if ca.NameSurname == "" {
			return errors.New("Requested Fields author..")
		}
		err = authorRepository.Create(ca)
		if err != nil {
			return err
		}

		autId, err := authorRepository.FindByAuthorName(ca.NameSurname)
		if err != nil {
			return err
		}
		ba := models.NewBook_Author(b.Book.ISBN, autId[0].ID)
		err = bookAuthorRepository.Create(*ba)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateAuthor(a AuthorItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	if a.NameSurname == "" {
		return errors.New("Requested Fields author..")
	}
	err := authorRepository.Create(a.Author)
	if err != nil {
		return err
	}
	autID, _ := authorRepository.FindByAuthorName(a.NameSurname)
	for _, b := range a.Book {
		// If we want, we can add other logical errors in if
		if b.ISBN <= 0 || b.Name == "" || b.StockAmount < 0 {
			return errors.New("Requested Fields name and isbn..")
		}
		err = bookRepository.Create(b)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		ba := models.NewBook_Author(b.ISBN, autID[0].ID)
		err = bookAuthorRepository.Create(*ba)
		if err != nil {
			return err
		}
	}

	return nil
}
