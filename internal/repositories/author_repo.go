package repositories

import (
	"errors"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/models"
	"gorm.io/gorm"
)

var (
	declaredAuthor = errors.New("Author has been declared")
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Migration() {
	r.db.AutoMigrate(&models.Author{})
}

// InsertSampleData adds datas from csv when program runs
func (r *AuthorRepository) InsertSampleData(authors []models.Author) {
	for _, c := range authors {
		r.Create(c)
	}
}

// GetAllAuthors takes all authors
func (r *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := r.db.Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//FindByAuthor finds Author with given author
func (r *AuthorRepository) FindByAuthor(author models.Author) (models.Author, error) {
	var athr models.Author
	result := r.db.Where(&models.Author{NameSurname: author.NameSurname}).First(&athr)

	if result.Error != nil {
		return author, result.Error
	}

	return athr, nil
}

//FindByAuthorName finds all Authors with given input
func (r *AuthorRepository) FindByAuthorName(input string) ([]models.Author, error) {
	var authors []models.Author
	result := r.db.Where("LOWER(name_surname) LIKE LOWER(?)", "%"+input+"%").Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//FindByID finds author with given id
func (r *AuthorRepository) FindByID(id int) ([]models.Author, error) {
	var authors []models.Author
	result := r.db.Find(&authors, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//Delete deletes author from base with parameter author
func (r *AuthorRepository) Delete(a models.Author) error {
	result := r.db.Delete(a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//DeleteByID deletes author from base with given id
func (r *AuthorRepository) DeleteByID(id int) error {
	result := r.db.Delete(&models.Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Create creats with given author
func (r *AuthorRepository) Create(author models.Author) error {
	result := r.db.Where(&models.Author{NameSurname: author.NameSurname}).First(&author)
	if result.Error != nil {
		r.db.Create(&author)
		return nil
	}

	return declaredAuthor
}

//Update updates with given author
func (r *AuthorRepository) Update(author models.Author) error {
	r.db.Save(author)
	return nil
}
