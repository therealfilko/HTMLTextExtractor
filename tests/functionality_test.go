package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    extractor "github.com/therealfilko/HTMLTextExtractor/pkg/extractor"
)

func TestFetchAndExtractHTML(t *testing.T) {
    // Mock-Server erstellen
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><div id="test">Hello, World!</div></body></html>`))
    }))
    defer mockServer.Close()

    // Testfall
    tests := []struct {
        name     string
        url      string
        selector string
        expected string
    }{
        {"ValidSelector", mockServer.URL, "#test", "Hello, World!"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := extractor.FetchAndExtractHTML(tt.url, tt.selector)
            if err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            if strings.TrimSpace(result) != tt.expected {
                t.Errorf("expected: %v, got: %v", tt.expected, result)
            }
        })
    }
}
