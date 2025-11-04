package ginrouter

import "testing"

func TestNewIDShouldBeIntError(t *testing.T) {
	err := NewIDShouldBeIntError("id")

	if err == nil {
		t.Fatal("Expected non-nil error")
	}

	expected := "bad request: id should be an integer"
	if err.Message != expected {
		t.Errorf("Expected message %s, got %s", expected, err.Message)
	}
}

func TestNewIDShouldBeIntError_DifferentParam(t *testing.T) {
	err := NewIDShouldBeIntError("userId")

	expected := "bad request: userId should be an integer"
	if err.Message != expected {
		t.Errorf("Expected message %s, got %s", expected, err.Message)
	}
}

func TestBadRequestError_StatusCode(t *testing.T) {
	err := &BadRequestError{Message: "test error"}

	statusCode := err.StatusCode()
	expected := 400

	if statusCode != expected {
		t.Errorf("Expected status code %d, got %d", expected, statusCode)
	}
}

func TestBadRequestError_Error(t *testing.T) {
	message := "test error message"
	err := &BadRequestError{Message: message}

	errorMsg := err.Error()

	if errorMsg != message {
		t.Errorf("Expected error message %s, got %s", message, errorMsg)
	}
}

func TestBadRequestError_ImplementsError(t *testing.T) {
	var _ error = &BadRequestError{}
}

func TestBadRequestError_ImplementsCustomError(t *testing.T) {
	var _ CustomError = &BadRequestError{}
}
