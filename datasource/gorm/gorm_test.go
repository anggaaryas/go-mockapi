package gormsql

import (
	"os"
	"testing"

	"github.com/anggaaryas/go-mockapi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	return db
}

func TestCreate(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	if ds == nil {
		t.Fatal("Expected Create to return a non-nil DataSource")
	}
}

func TestGetBaseURL_WithEnvVar(t *testing.T) {
	// Save original and set test value
	originalURL := os.Getenv("BASE_URL")
	testURL := "http://test.example.com"
	os.Setenv("BASE_URL", testURL)
	defer os.Setenv("BASE_URL", originalURL)

	result := getBaseURL()
	if result != testURL {
		t.Errorf("Expected base URL %s, got %s", testURL, result)
	}
}

func TestGetBaseURL_WithoutEnvVar(t *testing.T) {
	// Save original and clear
	originalURL := os.Getenv("BASE_URL")
	os.Unsetenv("BASE_URL")
	defer os.Setenv("BASE_URL", originalURL)

	result := getBaseURL()
	expected := "http://localhost:8080"
	if result != expected {
		t.Errorf("Expected base URL %s, got %s", expected, result)
	}
}

func TestGetCoverURL(t *testing.T) {
	originalURL := os.Getenv("BASE_URL")
	os.Setenv("BASE_URL", "http://test.com")
	defer os.Setenv("BASE_URL", originalURL)

	result := getCoverURL("test.jpg")
	expected := "http://test.com/static/image/test.jpg"
	if result != expected {
		t.Errorf("Expected cover URL %s, got %s", expected, result)
	}
}

func TestGetInitialBooks(t *testing.T) {
	books := getInitialBooks()

	if len(books) == 0 {
		t.Fatal("Expected getInitialBooks to return non-empty slice")
	}

	// Verify expected number of books
	expectedCount := 50
	if len(books) != expectedCount {
		t.Errorf("Expected %d books, got %d", expectedCount, len(books))
	}

	// Verify first book
	firstBook := books[0]
	if firstBook.ID != 1 {
		t.Errorf("Expected first book ID to be 1, got %d", firstBook.ID)
	}
	if firstBook.Title == "" {
		t.Error("Expected first book to have a title")
	}
	if firstBook.Author == "" {
		t.Error("Expected first book to have an author")
	}

	// Verify all books have required fields
	for i, book := range books {
		if book.ID == 0 {
			t.Errorf("Book at index %d has no ID", i)
		}
		if book.Title == "" {
			t.Errorf("Book at index %d has no title", i)
		}
		if book.Author == "" {
			t.Errorf("Book at index %d has no author", i)
		}
	}
}

func TestPopulateData_EmptyDatabase(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	// Verify books were created
	var count int64
	db.Model(&mockapi.Book{}).Count(&count)

	if count == 0 {
		t.Error("Expected books to be populated")
	}

	expectedCount := int64(50)
	if count != expectedCount {
		t.Errorf("Expected %d books, got %d", expectedCount, count)
	}
}

func TestPopulateData_AlreadyPopulated(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	// First population
	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("First PopulateData failed: %v", err)
	}

	var countAfterFirst int64
	db.Model(&mockapi.Book{}).Count(&countAfterFirst)

	// Second population should not add duplicates
	err = ds.PopulateData()
	if err != nil {
		t.Fatalf("Second PopulateData failed: %v", err)
	}

	var countAfterSecond int64
	db.Model(&mockapi.Book{}).Count(&countAfterSecond)

	if countAfterFirst != countAfterSecond {
		t.Errorf("Expected count to remain %d, got %d", countAfterFirst, countAfterSecond)
	}
}

func TestGetBookByID_Success(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	book, err := ds.GetBookByID("1")
	if err != nil {
		t.Fatalf("GetBookByID failed: %v", err)
	}

	if book.ID != 1 {
		t.Errorf("Expected book ID 1, got %d", book.ID)
	}
	if book.Title == "" {
		t.Error("Expected book to have a title")
	}
}

func TestGetBookByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	_, err = ds.GetBookByID("9999")
	if err == nil {
		t.Error("Expected error for non-existent book")
	}
}

func TestGetBooks_FirstPage(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	books, err := ds.GetBooks(1, 10, "")
	if err != nil {
		t.Fatalf("GetBooks failed: %v", err)
	}

	if len(books) != 10 {
		t.Errorf("Expected 10 books, got %d", len(books))
	}
}

func TestGetBooks_SecondPage(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	books, err := ds.GetBooks(2, 10, "")
	if err != nil {
		t.Fatalf("GetBooks failed: %v", err)
	}

	if len(books) != 10 {
		t.Errorf("Expected 10 books, got %d", len(books))
	}
}

func TestGetBooks_WithSearch_Title(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	books, err := ds.GetBooks(1, 10, "Go")
	if err != nil {
		t.Fatalf("GetBooks with search failed: %v", err)
	}

	if len(books) == 0 {
		t.Error("Expected to find books with 'Go' in title")
	}

	// Verify search result contains the search term
	found := false
	for _, book := range books {
		if contains(book.Title, "Go") || contains(book.Author, "Go") {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected at least one book to contain 'Go' in title or author")
	}
}

func TestGetBooks_WithSearch_Author(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	books, err := ds.GetBooks(1, 10, "Martin")
	if err != nil {
		t.Fatalf("GetBooks with author search failed: %v", err)
	}

	if len(books) == 0 {
		t.Error("Expected to find books with 'Martin' in author")
	}
}

func TestGetBooksCount_NoSearch(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	count, err := ds.GetBooksCount("")
	if err != nil {
		t.Fatalf("GetBooksCount failed: %v", err)
	}

	expectedCount := int64(50)
	if count != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, count)
	}
}

func TestGetBooksCount_WithSearch(t *testing.T) {
	db := setupTestDB(t)
	ds := Create(db)

	err := ds.PopulateData()
	if err != nil {
		t.Fatalf("PopulateData failed: %v", err)
	}

	count, err := ds.GetBooksCount("Go")
	if err != nil {
		t.Fatalf("GetBooksCount with search failed: %v", err)
	}

	if count == 0 {
		t.Error("Expected to find books with 'Go' in title or author")
	}

	// Count should be less than total
	totalCount, _ := ds.GetBooksCount("")
	if count >= totalCount {
		t.Errorf("Expected search count %d to be less than total %d", count, totalCount)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || contains(s[1:], substr)))
}
