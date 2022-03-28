package helpers

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/infrastructure"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/repositories"
)

const myPostgre_db = "host=localhost user=postgres password=pass1234 dbname=library port=5432 sslmode=disable" // arrange it for your db own connections

func ConnectDB() (*repositories.BookRepository, *repositories.AuthorRepository, *repositories.BookAuthRepository) {
	var bookRepository *repositories.BookRepository
	var authorRepository *repositories.AuthorRepository
	var bookAuthorRepository *repositories.BookAuthRepository

	db := infrastructure.NewPostgresDB(myPostgre_db)

	bookRepository = repositories.NewBookRepository(db)
	authorRepository = repositories.NewAuthorRepository(db)
	bookAuthorRepository = repositories.NewBookAuthRepository(db)

	bookRepository.Migration()
	authorRepository.Migration()
	bookAuthorRepository.Migration()

	bookRepository, authorRepository, bookAuthorRepository = ReadBookWithWorkerPool("../docs/books2.csv", bookRepository, authorRepository, bookAuthorRepository) //Take comment after fist run

	return bookRepository, authorRepository, bookAuthorRepository
}
