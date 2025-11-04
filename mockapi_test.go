package mockapi

import (
	"embed"
	"errors"
	"testing"
)

func TestGetStaticFiles(t *testing.T) {
	fs := GetStaticFiles()
	
	// Check that we get a valid embed.FS
	var _ embed.FS = fs
	
	// Try to read the static directory
	entries, err := fs.ReadDir("static")
	if err != nil {
		t.Fatalf("Expected to read static directory, got error: %v", err)
	}
	
	// Check that static directory has contents
	if len(entries) == 0 {
		t.Error("Expected static directory to have contents")
	}
}

type mockRouter struct {
	setupMockApiRouteFunc func(service Service) error
}

func (m *mockRouter) SetupMockApiRoute(service Service) error {
	if m.setupMockApiRouteFunc != nil {
		return m.setupMockApiRouteFunc(service)
	}
	return nil
}

func TestUse_Success(t *testing.T) {
	populateDataCalled := false
	setupRouteCalled := false
	
	ds := &mockDataSource{
		populateDataFunc: func() error {
			populateDataCalled = true
			return nil
		},
	}
	
	router := &mockRouter{
		setupMockApiRouteFunc: func(service Service) error {
			setupRouteCalled = true
			return nil
		},
	}
	
	// Use should not panic with successful operations
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Use panicked unexpectedly: %v", r)
		}
	}()
	
	Use(ds, router)
	
	if !populateDataCalled {
		t.Error("Expected PopulateData to be called")
	}
	
	if !setupRouteCalled {
		t.Error("Expected SetupMockApiRoute to be called")
	}
}

func TestUse_PopulateDataError(t *testing.T) {
	expectedError := errors.New("populate data error")
	
	ds := &mockDataSource{
		populateDataFunc: func() error {
			return expectedError
		},
	}
	
	router := &mockRouter{}
	
	// Use should panic when PopulateData fails
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Use to panic when PopulateData fails")
		} else if r != expectedError {
			t.Errorf("Expected panic with error %v, got %v", expectedError, r)
		}
	}()
	
	Use(ds, router)
}

func TestUse_SetupRouteError(t *testing.T) {
	expectedError := errors.New("setup route error")
	
	ds := &mockDataSource{
		populateDataFunc: func() error {
			return nil
		},
	}
	
	router := &mockRouter{
		setupMockApiRouteFunc: func(service Service) error {
			return expectedError
		},
	}
	
	// Use should panic when SetupMockApiRoute fails
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Use to panic when SetupMockApiRoute fails")
		} else if r != expectedError {
			t.Errorf("Expected panic with error %v, got %v", expectedError, r)
		}
	}()
	
	Use(ds, router)
}
