package models

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/pkg/funcs"
	"gorm.io/gorm"
)

//var id = 0 //for csv to slices after db adding; helper/readInsert

type Author struct {
	gorm.Model
	NameSurname string `json:"author"`
	Age         int    `json:"age"`
}

// func NewAuthor(nameSurname string) *Author { //for csv to slices after db adding; helper/readInsert
// 	id++
// 	return &Author{
// 		ID:          uint(id),
// 		NameSurname: nameSurname,
// 		Age:         funcs.RandomInt(20, 80),
// 	}
// }

func NewAuth(nameSurname string) *Author { //for csv to db adding ; helper/csvToDB
	return &Author{
		NameSurname: nameSurname,
		Age:         funcs.RandomInt(20, 80),
	}
}

//ToString func for Authors
func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name Surname: %s, Age: %d", a.ID, a.NameSurname, a.Age)
}

// BeforeDelete for taking info
func (a *Author) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Author (%s) deleting...", a.NameSurname)
	return nil
}
