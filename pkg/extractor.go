package extractor

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func FetchAndExtractHTML(url string, selector string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
    }

    // Load the HTML Document
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", err
    }

    result, err := doc.Find(selector).Html()
    if err != nil {
        return "", err
    }
    result_trimmed := strings.TrimSpace(result)
    return result_trimmed, nil
}
