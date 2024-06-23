package fetcher

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/therealfilko/HTMLTextExtractor/pkg/mock"
)

func TestFetchAndSaveHTML(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("mocked HTML content"))
    }))
    defer mockServer.Close()

    mockStorage := &mock.MockStorage{}
    err := FetchAndSaveHTML(mockServer.URL, mockStorage)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if mockStorage.Data != "mocked HTML content" {
        t.Errorf("expected 'mocked HTML content', got '%v'", mockStorage.Data)
    }
}
