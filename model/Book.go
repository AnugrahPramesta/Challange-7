package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// type Employee struct {
// 	ID       int    `json:"id" db:"id"`
// 	Fullname string `json:"full_name" db:"full_name"`
// 	Email    string `json:"email" db:"email"`
// 	Age      int    `json:"age" db:"age"`
// 	Division string `json:"division" db:"division"`
// }

type Books struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	NameBook  string    `json:"name_book" gorm:"unique;type:varchar(50)"`
	Author    string    `json:"author" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e Books) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.NameBook, validation.Required),
		validation.Field(&e.Author, validation.Required))
}
