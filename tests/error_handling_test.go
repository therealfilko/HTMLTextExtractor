package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    extractor "github.com/therealfilko/HTMLTextExtractor/pkg"
)

func TestFetchAndExtractHTMLExceptions(t *testing.T) {
    // Mock-Server für erfolgreichen Fall
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><div id="test">Hello, World!</div></body></html>`))
    }))
    defer mockServer.Close()

    // Mock-Server für Netzwerkfehler
    errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }))
    defer errorServer.Close()

    // Testfälle für Fehlerbedingungen
    tests := []struct {
        url      string
        selector string
        hasError bool
    }{
        {mockServer.URL, "#nonexistent", true},
        {errorServer.URL, "#test", true},
        {"http://invalid-url", "#test", true},
    }

    for _, tt := range tests {
        t.Run(tt.selector, func(t *testing.T) {
            _, err := extractor.FetchAndExtractHTML(tt.url, tt.selector)
            if (err != nil) != tt.hasError {
                t.Errorf("expected error: %v, got: %v", tt.hasError, err)
            }
        })
    }
}
