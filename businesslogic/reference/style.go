package reference

import (
	"time"
)

type Style struct {
	ID              int
	Name            string
	Description     string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStyleCriteria struct {
	StyleID int    `schema:"id"`
	Name    string `schema:"name"`
}

type IStyleRepository interface {
	CreateStyle(style *Style) error
	SearchStyle(criteria *SearchStyleCriteria) ([]Style, error)
	UpdateStyle(style Style) error
	DeleteStyle(style Style) error
}
