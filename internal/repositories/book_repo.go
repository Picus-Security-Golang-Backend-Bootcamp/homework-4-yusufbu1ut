package repositories

import (
	"errors"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/models"
	"gorm.io/gorm"
)

var (
	hihgerAmount = errors.New("Given count is higher than book' amount.")
	declaredBook = errors.New("Book has been declared")
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&models.Book{})
}

// InsertSampleData creates for db books when program runs
func (r *BookRepository) InsertSampleData(books []models.Book) {
	for _, c := range books {
		r.Create(c)
	}
}

// GetAllBooks gives all books
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// FindByBookName finds books with given input
func (r *BookRepository) FindByName(input string) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Where("LOWER(name) LIKE LOWER(?)", "%"+input+"%").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// FindByID finds book with parameter id
func (r *BookRepository) FindByID(id int) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

//FindByISBN finds books with ISBN uniqe count it is for book_authors base
func (r *BookRepository) FindByISBN(isbn int) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books, "isbn = ?", isbn)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// Deletes book with given book
func (r *BookRepository) Delete(b models.Book) error {
	_, err := r.FindByID(int(b.ID))
	if err != nil {
		return err
	}
	result := r.db.Delete(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Deletes book with given id
func (r *BookRepository) DeleteByID(id int) ([]models.Book, error) {
	book, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Delete(&models.Book{}, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

// Create creates book if it is not exist and uniqe looking on isbn name publisher
func (r *BookRepository) Create(book models.Book) error {
	result, _ := r.FindByISBN(book.ISBN)
	if len(result) != 0 {
		return declaredBook
	}
	r.db.Create(&book)
	return nil
}

// Update saves new changed book writing on db same id
func (r *BookRepository) Update(book models.Book) error {
	r.db.Save(book)
	return nil
}

// Buy decrease book stock amount with given count parameter
func (r *BookRepository) Buy(book models.Book, count int) error {

	if book.StockAmount < count {
		return hihgerAmount
	}

	book.StockAmount -= count

	r.Update(book)

	return nil
}
