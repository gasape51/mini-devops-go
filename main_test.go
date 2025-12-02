package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected to fail :
	expected := "Goodbye, World! "

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}
