package gormsql

import (
	_ "embed"
	"encoding/json"

	"github.com/anggaaryas/go-mockapi/mockapi"
	"gorm.io/gorm"
)

//go:embed books_data.json
var booksDataJSON []byte

type dataSource struct {
	db *gorm.DB
}

func getCoverURL(filename string) string {
	return "https://example.com/covers/" + filename
}

func Create(db *gorm.DB) mockapi.DataSource {
	return &dataSource{
		db: db,
	}
}

func getInitialBooks() []mockapi.Book {
	var books []mockapi.Book
	if err := json.Unmarshal(booksDataJSON, &books); err != nil {
		panic("failed to load books data: " + err.Error())
	}
	return books
}

func (ds *dataSource) PopulateData() error {
	err := ds.db.AutoMigrate(&mockapi.Book{})
	if err != nil {
		panic("failed to migrate database")
	}
	return ds.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&mockapi.Book{}).Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return nil
		}

		books := getInitialBooks()

		if err := tx.Create(&books).Error; err != nil {
			return err
		}

		return nil
	})
}

func (ds *dataSource) GetBookByID(id string) (mockapi.Book, error) {
	var book mockapi.Book
	if err := ds.db.First(&book, "id = ?", id).Error; err != nil {
		return mockapi.Book{}, err
	}
	return book, nil
}

func (ds *dataSource) GetBooks(page int, pageSize int) ([]mockapi.Book, error) {
	var books []mockapi.Book
	offset := (page - 1) * pageSize
	if err := ds.db.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
