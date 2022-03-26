package models

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/pkg/funcs"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	NameSurname string `json:"author"`
	Age         int    `json:"age"`
}

func NewAuth(nameSurname string) *Author {
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
