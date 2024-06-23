package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    extractor "github.com/therealfilko/HTMLTextExtractor/pkg/extractor"
)

func TestFetchAndExtractHTMLErrorHandling(t *testing.T) {
    // Mock-Server für Netzwerkfehler
    errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }))
    defer errorServer.Close()

    // Testfälle für Fehlerbedingungen
    tests := []struct {
        name     string
        url      string
        selector string
        hasError bool
    }{
        {"NonexistentSelector", errorServer.URL, "#nonexistent", true},
        {"HTTPError", errorServer.URL, "#test", true},
        {"InvalidURL", "http://invalid-url", "#test", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := extractor.FetchAndExtractHTML(tt.url, tt.selector)
            if (err != nil) != tt.hasError {
                t.Errorf("expected error: %v, got: %v", tt.hasError, err)
            }
        })
    }
}
