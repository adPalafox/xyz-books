package dao

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              int    `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	Isbn10          *string
	Isbn13          *string
	ListPrice       float32   `gorm:"not null"`
	PublicationYear int       `gorm:"not null"`
	PublisherID     int       `gorm:"not null"`
	Publisher       Publisher `gorm:"foreignKey:PublisherID"`
	ImageUrl        *string
	Edition         *string
	Authors         []Author  `gorm:"many2many:book_authors;"`
	CreatedAt       time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type Author struct {
	ID         int    `gorm:"primaryKey"`
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	MiddleName *string
	Books      []Book    `gorm:"many2many:book_authors;"`
	CreatedAt  time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type Publisher struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Books     []Book    `gorm:"many2one:books"`
	CreatedAt time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
