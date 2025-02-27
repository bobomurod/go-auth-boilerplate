package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHomeHandler(t *testing.T) {
	r := setupRouter()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("http.NewRequest failed: %v", err)
		return
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}
	expectedBody := "Hello, World!"
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("error reading body: %v", err)
	}
	if string(body) != expectedBody {
		t.Fatalf("handler returned unexpected body: got %q want %q", string(body), expectedBody)
	}
}

func testNonExistingRoute(t *testing.T) {
	r := setupRouter()
	req, err := http.NewRequest("GET", "/non-existing-route", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusNotFound)
	}
}
