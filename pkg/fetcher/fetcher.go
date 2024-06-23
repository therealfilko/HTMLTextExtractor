package fetcher

import (
    "io"
    "net/http"
    "github.com/therealfilko/HTMLTextExtractor/pkg/storage"
)

func FetchAndSaveHTML(url string, s storage.Storage) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    return s.Save(string(body))
}
