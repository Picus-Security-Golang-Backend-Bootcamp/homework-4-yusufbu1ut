package helpers

import (
	"net/http"

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

// ListBook returns books with author in BookItem slice
func ListBook(bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem
	books, err := bookRepository.GetAllBooks()
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	for _, b := range books {
		var item BookItem
		item.Book = b
		book_authors, err := bookAuthorRepository.FindByISBN(b.ISBN)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		for _, ba := range book_authors {
			author, err := authorRepository.FindByID(int(ba.AuthorID))
			if err != nil {
				msg := err.Error()
				res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
				return nil, res
			}
			item.Auth = append(item.Auth, author[0])
		}
		items = append(items, item)
	}
	return items, nil
}

// ListAuth returns authors with books in AuthorItem slice
func ListAuth(bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem
	authors, err := authorRepository.GetAllAuthors()
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	for _, a := range authors {
		var item AuthorItem
		item.Author = a
		book_authors, err := bookAuthorRepository.FindByAuthorID(int(a.ID))
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		for _, ba := range book_authors {
			book, err := bookRepository.FindByISBN(ba.BookID)
			if err != nil {
				msg := err.Error()
				res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
				return nil, res
			}
			item.Book = append(item.Book, book[0])
		}
		items = append(items, item)
	}
	return items, nil
}

// SearchByBookInput takes input parameter and checks in books and returns with BookItem slice
func SearchByBookInput(srch string, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem
	books, err := bookRepository.FindByName(srch)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	if len(books) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return nil, res
	}
	for _, b := range books {
		var item BookItem
		item.Book = b
		book_authors, err := bookAuthorRepository.FindByISBN(b.ISBN)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		for _, ba := range book_authors {
			author, err := authorRepository.FindByID(int(ba.AuthorID))
			if err != nil {
				msg := err.Error()
				res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
				return nil, res
			}
			item.Auth = append(item.Auth, author[0])
		}
		items = append(items, item)
	}
	return items, nil
}

// SearchByAuthorInput takes input parameter and checks in authors and returns with AuthorItem slice
func SearchByAuthorInput(srch string, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem
	authors, err := authorRepository.FindByAuthorName(srch)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	if len(authors) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return nil, res
	}
	for _, a := range authors {
		var item AuthorItem
		item.Author = a
		book_authors, err := bookAuthorRepository.FindByAuthorID(int(a.ID))
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		for _, ba := range book_authors {
			book, err := bookRepository.FindByISBN(ba.BookID)
			if err != nil {
				msg := err.Error()
				res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
				return nil, res
			}
			item.Book = append(item.Book, book[0])
		}
		items = append(items, item)
	}
	return items, nil
}

// DeleteByBookID takes int count and deletes which as connected bases in books and book_authors
func DeleteByBookID(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	book, err := bookRepository.DeleteByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	if len(book) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return res
	}
	bookAuthorRepository.DeleteByISBN(book[0].ISBN)

	return nil
}

// DeleteByAuthorID takes int count and deletes which as connected bases in authors and book_authors
func DeleteByAuthorID(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	auth, err := authorRepository.FindByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusAlreadyReported, Msg: msg}
		return res
	}
	if len(auth) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return res
	}
	err = authorRepository.DeleteByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	bookAuthorRepository.DeleteByAuthorID(id)

	return nil
}

// GetBookById gets Book with given id and reponse with BookItem
func GetBookById(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]BookItem, error) {
	var items []BookItem
	var item BookItem
	book, err := bookRepository.FindByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	if len(book) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return nil, res
	}
	item.Book = book[0]
	ba, err := bookAuthorRepository.FindByISBN(book[0].ISBN)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	for _, a := range ba {
		author, err := authorRepository.FindByID(int(a.AuthorID))
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		item.Auth = append(item.Auth, author[0])
	}
	items = append(items, item)
	return items, nil
}

// GetAuthById gets Author with given id and reponse with AuthorItem
func GetAuthById(id int, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) ([]AuthorItem, error) {
	var items []AuthorItem
	var item AuthorItem
	auth, err := authorRepository.FindByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	if len(auth) == 0 {
		msg := "No Data"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return nil, res
	}
	item.Author = auth[0]
	ba, err := bookAuthorRepository.FindByAuthorID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return nil, res
	}
	for _, b := range ba {
		book, err := bookRepository.FindByISBN(b.BookID)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return nil, res
		}
		item.Book = append(item.Book, book[0])
	}
	items = append(items, item)
	return items, nil
}

// CreateBook takes a BookItem and creates Book, name,isbn,amount fields should given
// 	if it has authors in body field just making connections on book_author
func CreateBook(b BookItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	// If we want, we can add other logical errors in if
	if b.ISBN <= 0 || b.Name == "" || b.StockAmount < 0 {
		msg := "Requested Fields.."
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	err := bookRepository.Create(b.Book)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	for _, ca := range b.Auth {
		if ca.NameSurname == "" {
			msg := "Requested Fields.."
			res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
			return res
		}
		err = authorRepository.Create(ca)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}

		autId, err := authorRepository.FindByAuthorName(ca.NameSurname)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
		ba := models.NewBook_Author(b.Book.ISBN, autId[0].ID)
		err = bookAuthorRepository.Create(*ba)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
	}
	return nil
}

// CreateAuthor takes an AuthorItem and creates Author, namesurname field should given
// 	if it has books in body field just making connections on book_author
func CreateAuthor(a AuthorItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	if a.NameSurname == "" {
		msg := "Requested Fields.."
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	err := authorRepository.Create(a.Author)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	autID, err := authorRepository.FindByAuthorName(a.NameSurname)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	for _, b := range a.Book {
		// If we want, we can add other logical errors in if
		if b.ISBN <= 0 || b.Name == "" || b.StockAmount < 0 {
			msg := "Requested Fields.."
			res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
			return res
		}
		err = bookRepository.Create(b)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
		ba := models.NewBook_Author(b.ISBN, autID[0].ID)
		err = bookAuthorRepository.Create(*ba)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
	}

	return nil
}

// UpdateBook takes a BookItem from request body and checks it if it is avaiable changes and updates data
//	request should must have and book id, book name,book amount if body has an author it should have a namesurname
// 	in here updates works on name,price,publisher,pages,amount these information should be given if dont then empty data replace
// 	if it has authors in body field just making connections on book_author
func UpdateBook(b BookItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	if b.Name == "" || b.StockAmount < 0 || b.ID <= 0 {
		msg := "Requested Fields amount,name and isbn.."
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	book, err := bookRepository.FindByID(int(b.ID))
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	if len(book) == 0 {
		msg := "No data foun with given id"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return res
	}

	book[0].Name = b.Name
	book[0].Pages = b.Pages
	book[0].Price = b.Price
	book[0].Publisher = b.Publisher
	book[0].StockAmount = b.StockAmount
	bookRepository.Update(book[0])

	for _, ua := range b.Auth {
		if ua.NameSurname == "" {
			msg := "Requested fields for author.."
			res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
			return res
		}
		ath, err := authorRepository.FindByAuthorName(ua.NameSurname)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
		ba := models.NewBook_Author(book[0].ISBN, ath[0].ID)
		bookAuthorRepository.Create(*ba)
	}

	return nil
}

// UpdateAuthor takes an AuthorItem from request body and checks it if it is avaiable changes and updates data
//	request should must have and author id and author namesurname if body has a book it should have a name
// 	if it has books in body field just making connections on book_author
func UpdateAuthor(a AuthorItem, bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, bookAuthorRepository repositories.BookAuthRepository) error {
	if a.NameSurname == "" || a.ID <= 0 {
		msg := "Requested fields"
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	author, err := authorRepository.FindByID(int(a.ID))
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	if len(author) == 0 {
		msg := "No data foun with given id"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return res
	}

	author[0].NameSurname = a.NameSurname
	author[0].Age = a.Age
	authorRepository.Update(author[0])

	for _, ub := range a.Book {
		if ub.Name == "" {
			msg := "Requested Fields.."
			res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
			return res
		}
		book, err := bookRepository.FindByName(ub.Name)
		if err != nil {
			msg := err.Error()
			res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
			return res
		}
		ba := models.NewBook_Author(book[0].ISBN, a.ID)
		bookAuthorRepository.Create(*ba)
	}

	return nil
}

// BookBuy works on books and takes id removes on its amount and saves it
func BookBuy(id int, cnt int, bookRepository repositories.BookRepository) error {
	book, err := bookRepository.FindByID(id)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadGateway, Msg: msg}
		return res
	}
	if len(book) == 0 {
		msg := "No data foun with given id"
		res := &MalformedRequest{Status: http.StatusNotFound, Msg: msg}
		return res
	}
	err = bookRepository.Buy(book[0], cnt)
	if err != nil {
		msg := err.Error()
		res := &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		return res
	}
	return nil
}
