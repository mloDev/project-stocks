package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndex(t *testing.T) {
	projectID := os.Getenv("GOLANG_SAMPLES_FIRESTORE_PROJECT")
	if projectID == "" {
		t.Skip("Skipping Firestore test. Set GOLANG_SAMPLES_FIRESTORE_PROJECT.")
	}

	a, err := newApp(projectID, "")
	if err != nil {
		t.Fatalf("newApp: %v", err)
	}

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	a.index(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code, got %v, want %v", resp.StatusCode, http.StatusOK)
	}
}