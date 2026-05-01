package backend

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	router := SetupRouter()
	req := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["status"] != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response["status"])
	}
}

func TestCORSHeaders(t *testing.T) {
	router := SetupRouter()
	req := httptest.NewRequest("OPTIONS", "/api/health", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	allowOrigin := w.Header().Get("Access-Control-Allow-Origin")
	if allowOrigin != "http://localhost:5173" {
		t.Errorf("Expected 'http://localhost:5173', got '%s'", allowOrigin)
	}
}
