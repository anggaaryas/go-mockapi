package mockapi

import (
	"encoding/json"
	"testing"
)

func TestBook_JSONMarshaling(t *testing.T) {
	book := Book{
		ID:       1,
		Title:    "Test Book",
		Author:   "Test Author",
		Category: "Test Category",
		Desc:     "Test Description",
		CoverURL: "http://test.com/cover.jpg",
	}

	// Test marshaling
	data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal Book: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Book
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Book: %v", err)
	}

	// Verify all fields
	if unmarshaled.ID != book.ID {
		t.Errorf("Expected ID %d, got %d", book.ID, unmarshaled.ID)
	}
	if unmarshaled.Title != book.Title {
		t.Errorf("Expected title %s, got %s", book.Title, unmarshaled.Title)
	}
	if unmarshaled.Author != book.Author {
		t.Errorf("Expected author %s, got %s", book.Author, unmarshaled.Author)
	}
	if unmarshaled.Category != book.Category {
		t.Errorf("Expected category %s, got %s", book.Category, unmarshaled.Category)
	}
	if unmarshaled.Desc != book.Desc {
		t.Errorf("Expected desc %s, got %s", book.Desc, unmarshaled.Desc)
	}
	if unmarshaled.CoverURL != book.CoverURL {
		t.Errorf("Expected cover URL %s, got %s", book.CoverURL, unmarshaled.CoverURL)
	}
}

func TestBook_EmptyFields(t *testing.T) {
	book := Book{
		ID: 1,
	}

	data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal Book with empty fields: %v", err)
	}

	var unmarshaled Book
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Book with empty fields: %v", err)
	}

	if unmarshaled.ID != book.ID {
		t.Errorf("Expected ID %d, got %d", book.ID, unmarshaled.ID)
	}
	if unmarshaled.Title != "" {
		t.Errorf("Expected empty title, got %s", unmarshaled.Title)
	}
}

func TestBook_JSONTags(t *testing.T) {
	book := Book{
		ID:       1,
		Title:    "Test",
		Author:   "Author",
		Category: "Category",
		Desc:     "Description",
		CoverURL: "url",
	}

	data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal Book: %v", err)
	}

	jsonStr := string(data)

	// Check that JSON uses correct field names
	expectedFields := []string{"id", "title", "author", "category", "desc", "cover_url"}
	for _, field := range expectedFields {
		if !contains(jsonStr, field) {
			t.Errorf("Expected JSON to contain field %s", field)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || contains(s[1:], substr)))
}
