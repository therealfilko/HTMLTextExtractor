package main

import (
    "os"
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/therealfilko/HTMLTextExtractor/pkg/mock"
    "github.com/therealfilko/HTMLTextExtractor/pkg/fetcher"
)

func TestMainExtract(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><h1 id="test">Hello, World!</h1></body></html>`))
    }))
    defer mockServer.Close()

    os.Args = []string{"cmd", "extract", mockServer.URL, "#test"}
    main()
}

func TestMainCount(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`<html><body><h1 id="test">Hello, World! This is a test.</h1></body></html>`))
    }))
    defer mockServer.Close()

    os.Args = []string{"cmd", "count", mockServer.URL}
    main()
}

func TestMainFetchAndSave(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("mocked HTML content"))
    }))
    defer mockServer.Close()

    mockStorage := &mock.MockStorage{}
    fetcher.FetchAndSaveHTML(mockServer.URL, mockStorage)

    if mockStorage.Data != "mocked HTML content" {
        t.Errorf("expected 'mocked HTML content', got '%v'", mockStorage.Data)
    }
}
