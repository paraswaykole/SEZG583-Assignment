package books

import (
	"book-catalogue-service/src/db"
	"time"
)

type Book struct {
	ID              uint      `gorm:"primaryKey" json:"id"`             // ID field (primary key)
	Title           string    `gorm:"size:255;not null" json:"title"`   // Title of the book
	Author          string    `gorm:"size:255" json:"author"`           // Author of the book
	ISBN            string    `gorm:"size:13;unique" json:"isbn"`       // ISBN (unique identifier for the book)
	PublicationYear int       `gorm:"type:int" json:"publication_year"` // Year of publication
	Genre           string    `gorm:"size:50" json:"genre"`             // Genre of the book
	Summary         string    `gorm:"type:text" json:"summary"`         // Summary of the book
	CoverImageURL   string    `gorm:"size:255" json:"cover_image_url"`  // URL of the book's cover image
	Publisher       string    `gorm:"size:255" json:"publisher"`        // Publisher of the book
	CreatedAt       time.Time `json:"created_at"`                       // Creation timestamp
	UpdatedAt       time.Time `json:"updated_at"`                       // Update timestamp
}

func CreateBook(book *Book) (*Book, error) {
	if err := db.Get().Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBooks(offset, limit int, searchQuery string) ([]Book, error) {
	var books []Book
	query := db.Get().Model(&Book{}).Where("title LIKE ? OR author LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	query.Offset(offset).Limit(limit).Find(&books)
	return books, nil
}

func GetBookByID(id uint) (*Book, error) {
	var book Book
	if err := db.Get().First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func DeleteBook(id uint) error {
	if err := db.Get().Where("id = ?", id).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}
