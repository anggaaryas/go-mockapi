package ginrouter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anggaaryas/go-mockapi"
	"github.com/gin-gonic/gin"
)

type mockService struct {
	getBookByIDFunc func(id string) (mockapi.Book, error)
	getBooksFunc    func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error)
}

func (m *mockService) GetBookByID(id string) (mockapi.Book, error) {
	if m.getBookByIDFunc != nil {
		return m.getBookByIDFunc(id)
	}
	return mockapi.Book{}, nil
}

func (m *mockService) GetBooks(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
	if m.getBooksFunc != nil {
		return m.getBooksFunc(page, pageSize, search)
	}
	return mockapi.PaginatedBooks{}, nil
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestCreate(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	if router == nil {
		t.Fatal("Expected Create to return a non-nil Router")
	}
}

func TestSetupMockApiRoute(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{}
	err := router.SetupMockApiRoute(service)

	if err != nil {
		t.Fatalf("SetupMockApiRoute failed: %v", err)
	}
}

func TestGetBookByID_Success(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	expectedBook := mockapi.Book{
		ID:       1,
		Title:    "Test Book",
		Author:   "Test Author",
		Category: "Test Category",
		Desc:     "Test Description",
		CoverURL: "http://test.com/cover.jpg",
	}

	service := &mockService{
		getBookByIDFunc: func(id string) (mockapi.Book, error) {
			if id == "1" {
				return expectedBook, nil
			}
			return mockapi.Book{}, errors.New("not found")
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var book mockapi.Book
	err := json.Unmarshal(w.Body.Bytes(), &book)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if book.ID != expectedBook.ID {
		t.Errorf("Expected book ID %d, got %d", expectedBook.ID, book.ID)
	}
	if book.Title != expectedBook.Title {
		t.Errorf("Expected book title %s, got %s", expectedBook.Title, book.Title)
	}
}

func TestGetBookByID_InvalidID(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{}
	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var apiErr mockapi.APIError
	err := json.Unmarshal(w.Body.Bytes(), &apiErr)
	if err != nil {
		t.Fatalf("Failed to unmarshal error response: %v", err)
	}

	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code in error %d, got %d", http.StatusBadRequest, apiErr.StatusCode)
	}
}

func TestGetBookByID_NotFound(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{
		getBookByIDFunc: func(id string) (mockapi.Book, error) {
			return mockapi.Book{}, errors.New("book not found")
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestGetBooks_Success(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	expectedBooks := mockapi.PaginatedBooks{
		Data: []mockapi.Book{
			{ID: 1, Title: "Book 1", Author: "Author 1"},
			{ID: 2, Title: "Book 2", Author: "Author 2"},
		},
		Page:       1,
		PageSize:   10,
		TotalItems: 2,
		TotalPages: 1,
	}

	service := &mockService{
		getBooksFunc: func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
			return expectedBooks, nil
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var books mockapi.PaginatedBooks
	err := json.Unmarshal(w.Body.Bytes(), &books)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(books.Data) != 2 {
		t.Errorf("Expected 2 books, got %d", len(books.Data))
	}
}

func TestGetBooks_WithPagination(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{
		getBooksFunc: func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
			if page != 2 || pageSize != 5 {
				t.Errorf("Expected page 2 and pageSize 5, got page %d and pageSize %d", page, pageSize)
			}
			return mockapi.PaginatedBooks{
				Data:       []mockapi.Book{},
				Page:       page,
				PageSize:   pageSize,
				TotalItems: 20,
				TotalPages: 4,
			}, nil
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books?page=2&page_size=5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetBooks_WithSearch(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{
		getBooksFunc: func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
			if search != "Go" {
				t.Errorf("Expected search term 'Go', got %s", search)
			}
			return mockapi.PaginatedBooks{
				Data:       []mockapi.Book{{ID: 1, Title: "Go Programming"}},
				Page:       1,
				PageSize:   10,
				TotalItems: 1,
				TotalPages: 1,
			}, nil
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books?search=Go", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetBooks_DefaultParameters(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{
		getBooksFunc: func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
			if page != 1 {
				t.Errorf("Expected default page 1, got %d", page)
			}
			if pageSize != 10 {
				t.Errorf("Expected default pageSize 10, got %d", pageSize)
			}
			if search != "" {
				t.Errorf("Expected empty search, got %s", search)
			}
			return mockapi.PaginatedBooks{}, nil
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetBooks_Error(t *testing.T) {
	r := setupTestRouter()
	router := Create(r)

	service := &mockService{
		getBooksFunc: func(page int, pageSize int, search string) (mockapi.PaginatedBooks, error) {
			return mockapi.PaginatedBooks{}, errors.New("database error")
		},
	}

	router.SetupMockApiRoute(service)

	req, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestGetErrorResponse_CustomError(t *testing.T) {
	r := setupTestRouter()
	cfg := &config{r: r}

	customErr := NewIDShouldBeIntError("id")
	apiErr := cfg.getErrorResponse(customErr)

	if apiErr.StatusCode != 400 {
		t.Errorf("Expected status code 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != customErr.Message {
		t.Errorf("Expected message %s, got %s", customErr.Message, apiErr.Message)
	}
}

func TestGetErrorResponse_GenericError_TestMode(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	cfg := &config{r: r}

	testErr := errors.New("test error")
	apiErr := cfg.getErrorResponse(testErr)

	if apiErr.StatusCode != 500 {
		t.Errorf("Expected status code 500, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != testErr.Error() {
		t.Errorf("Expected message %s, got %s", testErr.Error(), apiErr.Message)
	}
}

func TestGetErrorResponse_GenericError_ReleaseMode(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	defer gin.SetMode(gin.TestMode)

	r := gin.New()
	cfg := &config{r: r}

	testErr := errors.New("test error")
	apiErr := cfg.getErrorResponse(testErr)

	if apiErr.StatusCode != 500 {
		t.Errorf("Expected status code 500, got %d", apiErr.StatusCode)
	}
	expectedMessage := "An error occurred while processing your request"
	if apiErr.Message != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, apiErr.Message)
	}
}
