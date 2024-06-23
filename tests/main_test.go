package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    extractor "github.com/therealfilko/HTMLTextExtractor/pkg"
)

func TestFetchAndExtractHTML(t *testing.T) {
    // Setup a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><div id="test">Hello, World!</div></body></html>`))
    }))
    defer server.Close()

    // Test cases
    tests := []struct {
        url      string
        selector string
        expected string
        hasError bool
    }{
        {server.URL, "#test", "Hello, World!", false},
        {server.URL, "#nonexistent", "", true},
        {"http://invalid-url", "#test", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.selector, func(t *testing.T) {
            result, err := extractor.FetchAndExtractHTML(tt.url, tt.selector)
            if (err != nil) != tt.hasError {
                t.Errorf("expected error: %v, got: %v", tt.hasError, err)
            }
            if strings.TrimSpace(result) != tt.expected {
                t.Errorf("expected: %v, got: %v", tt.expected, result)
            }
        })
    }
}
