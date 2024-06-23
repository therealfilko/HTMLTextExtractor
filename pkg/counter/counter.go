package counter

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func CountWords(url string) (int, error) {
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return 0, fmt.Errorf("Fehler: Status Code %d %s", resp.StatusCode, resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return 0, err
    }

    words := strings.Fields(doc.Text())
    return len(words), nil
}
