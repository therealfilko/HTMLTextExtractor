package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    counter "github.com/therealfilko/HTMLTextExtractor/pkg/counter"
)

func TestCountWords(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><p>Hello, World! This is a test.</p></body></html>`))
    }))
    defer mockServer.Close()

    result, err := counter.CountWords(mockServer.URL)
    if err != nil {
        t.Errorf("unerwarteter Fehler: %v", err)
    }
    if result != 6 {
        t.Errorf("erwartet: 6, erhalten: %v", result)
    }
}
