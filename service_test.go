package mockapi

import (
	"errors"
	"testing"
)

type mockDataSource struct {
	populateDataFunc  func() error
	getBookByIDFunc   func(id string) (Book, error)
	getBooksFunc      func(page int, pageSize int, search string) ([]Book, error)
	getBooksCountFunc func(search string) (int64, error)
}

func (m *mockDataSource) PopulateData() error {
	if m.populateDataFunc != nil {
		return m.populateDataFunc()
	}
	return nil
}

func (m *mockDataSource) GetBookByID(id string) (Book, error) {
	if m.getBookByIDFunc != nil {
		return m.getBookByIDFunc(id)
	}
	return Book{}, nil
}

func (m *mockDataSource) GetBooks(page int, pageSize int, search string) ([]Book, error) {
	if m.getBooksFunc != nil {
		return m.getBooksFunc(page, pageSize, search)
	}
	return []Book{}, nil
}

func (m *mockDataSource) GetBooksCount(search string) (int64, error) {
	if m.getBooksCountFunc != nil {
		return m.getBooksCountFunc(search)
	}
	return 0, nil
}

func TestNewService(t *testing.T) {
	ds := &mockDataSource{}
	svc := NewService(ds)

	if svc == nil {
		t.Fatal("NewService should return a non-nil service")
	}
}

func TestService_GetBookByID_Success(t *testing.T) {
	expectedBook := Book{
		ID:       1,
		Title:    "Test Book",
		Author:   "Test Author",
		Category: "Test Category",
		Desc:     "Test Description",
		CoverURL: "http://test.com/cover.jpg",
	}

	ds := &mockDataSource{
		getBookByIDFunc: func(id string) (Book, error) {
			if id == "1" {
				return expectedBook, nil
			}
			return Book{}, errors.New("book not found")
		},
	}

	svc := NewService(ds)
	book, err := svc.GetBookByID("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if book.ID != expectedBook.ID {
		t.Errorf("Expected book ID %d, got %d", expectedBook.ID, book.ID)
	}
	if book.Title != expectedBook.Title {
		t.Errorf("Expected book title %s, got %s", expectedBook.Title, book.Title)
	}
}

func TestService_GetBookByID_Error(t *testing.T) {
	expectedError := errors.New("database error")
	ds := &mockDataSource{
		getBookByIDFunc: func(id string) (Book, error) {
			return Book{}, expectedError
		},
	}

	svc := NewService(ds)
	_, err := svc.GetBookByID("1")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestService_GetBooks_Success(t *testing.T) {
	expectedBooks := []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1"},
		{ID: 2, Title: "Book 2", Author: "Author 2"},
	}

	ds := &mockDataSource{
		getBooksFunc: func(page int, pageSize int, search string) ([]Book, error) {
			return expectedBooks, nil
		},
		getBooksCountFunc: func(search string) (int64, error) {
			return 2, nil
		},
	}

	svc := NewService(ds)
	result, err := svc.GetBooks(1, 10, "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Data) != 2 {
		t.Errorf("Expected 2 books, got %d", len(result.Data))
	}
	if result.TotalItems != 2 {
		t.Errorf("Expected total items 2, got %d", result.TotalItems)
	}
	if result.Page != 1 {
		t.Errorf("Expected page 1, got %d", result.Page)
	}
	if result.PageSize != 10 {
		t.Errorf("Expected page size 10, got %d", result.PageSize)
	}
	if result.TotalPages != 1 {
		t.Errorf("Expected total pages 1, got %d", result.TotalPages)
	}
}

func TestService_GetBooks_WithSearch(t *testing.T) {
	expectedBooks := []Book{
		{ID: 1, Title: "Go Programming", Author: "Author 1"},
	}

	ds := &mockDataSource{
		getBooksFunc: func(page int, pageSize int, search string) ([]Book, error) {
			if search == "Go" {
				return expectedBooks, nil
			}
			return []Book{}, nil
		},
		getBooksCountFunc: func(search string) (int64, error) {
			if search == "Go" {
				return 1, nil
			}
			return 0, nil
		},
	}

	svc := NewService(ds)
	result, err := svc.GetBooks(1, 10, "Go")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Data) != 1 {
		t.Errorf("Expected 1 book, got %d", len(result.Data))
	}
	if result.TotalItems != 1 {
		t.Errorf("Expected total items 1, got %d", result.TotalItems)
	}
}

func TestService_GetBooks_Pagination(t *testing.T) {
	ds := &mockDataSource{
		getBooksFunc: func(page int, pageSize int, search string) ([]Book, error) {
			return []Book{}, nil
		},
		getBooksCountFunc: func(search string) (int64, error) {
			return 25, nil
		},
	}

	svc := NewService(ds)
	result, err := svc.GetBooks(1, 10, "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.TotalPages != 3 {
		t.Errorf("Expected total pages 3, got %d", result.TotalPages)
	}
}

func TestService_GetBooks_Error(t *testing.T) {
	expectedError := errors.New("database error")
	ds := &mockDataSource{
		getBooksFunc: func(page int, pageSize int, search string) ([]Book, error) {
			return []Book{}, expectedError
		},
		getBooksCountFunc: func(search string) (int64, error) {
			return 0, expectedError
		},
	}

	svc := NewService(ds)
	_, err := svc.GetBooks(1, 10, "")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
