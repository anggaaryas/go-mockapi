package mockapi

import (
	"encoding/json"
	"testing"
)

func TestPaginatedBooks_JSONMarshaling(t *testing.T) {
	books := []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Category: "Cat 1", Desc: "Desc 1", CoverURL: "url1"},
		{ID: 2, Title: "Book 2", Author: "Author 2", Category: "Cat 2", Desc: "Desc 2", CoverURL: "url2"},
	}
	
	pb := PaginatedBooks{
		Data:       books,
		Page:       1,
		PageSize:   10,
		TotalItems: 25,
		TotalPages: 3,
	}
	
	// Test marshaling
	data, err := json.Marshal(pb)
	if err != nil {
		t.Fatalf("Failed to marshal PaginatedBooks: %v", err)
	}
	
	// Test unmarshaling
	var unmarshaled PaginatedBooks
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal PaginatedBooks: %v", err)
	}
	
	// Verify fields
	if unmarshaled.Page != pb.Page {
		t.Errorf("Expected page %d, got %d", pb.Page, unmarshaled.Page)
	}
	if unmarshaled.PageSize != pb.PageSize {
		t.Errorf("Expected page size %d, got %d", pb.PageSize, unmarshaled.PageSize)
	}
	if unmarshaled.TotalItems != pb.TotalItems {
		t.Errorf("Expected total items %d, got %d", pb.TotalItems, unmarshaled.TotalItems)
	}
	if unmarshaled.TotalPages != pb.TotalPages {
		t.Errorf("Expected total pages %d, got %d", pb.TotalPages, unmarshaled.TotalPages)
	}
	if len(unmarshaled.Data) != len(pb.Data) {
		t.Errorf("Expected %d books, got %d", len(pb.Data), len(unmarshaled.Data))
	}
}

func TestPaginatedBooks_EmptyData(t *testing.T) {
	pb := PaginatedBooks{
		Data:       []Book{},
		Page:       1,
		PageSize:   10,
		TotalItems: 0,
		TotalPages: 0,
	}
	
	data, err := json.Marshal(pb)
	if err != nil {
		t.Fatalf("Failed to marshal empty PaginatedBooks: %v", err)
	}
	
	var unmarshaled PaginatedBooks
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal empty PaginatedBooks: %v", err)
	}
	
	if len(unmarshaled.Data) != 0 {
		t.Errorf("Expected 0 books, got %d", len(unmarshaled.Data))
	}
}

func TestAPIError_JSONMarshaling(t *testing.T) {
	apiErr := APIError{
		StatusCode: 404,
		Message:    "Book not found",
	}
	
	// Test marshaling
	data, err := json.Marshal(apiErr)
	if err != nil {
		t.Fatalf("Failed to marshal APIError: %v", err)
	}
	
	// Test unmarshaling
	var unmarshaled APIError
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal APIError: %v", err)
	}
	
	// Verify fields
	if unmarshaled.StatusCode != apiErr.StatusCode {
		t.Errorf("Expected status code %d, got %d", apiErr.StatusCode, unmarshaled.StatusCode)
	}
	if unmarshaled.Message != apiErr.Message {
		t.Errorf("Expected message %s, got %s", apiErr.Message, unmarshaled.Message)
	}
}

func TestAPIError_DifferentStatusCodes(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		message    string
	}{
		{"BadRequest", 400, "Bad request"},
		{"NotFound", 404, "Not found"},
		{"InternalError", 500, "Internal server error"},
		{"Unauthorized", 401, "Unauthorized"},
		{"Forbidden", 403, "Forbidden"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiErr := APIError{
				StatusCode: tc.statusCode,
				Message:    tc.message,
			}
			
			if apiErr.StatusCode != tc.statusCode {
				t.Errorf("Expected status code %d, got %d", tc.statusCode, apiErr.StatusCode)
			}
			if apiErr.Message != tc.message {
				t.Errorf("Expected message %s, got %s", tc.message, apiErr.Message)
			}
		})
	}
}
