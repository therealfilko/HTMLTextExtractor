package extractor

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// FetchAndExtractHTML fetches the URL and extracts HTML based on the selector
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

    result := doc.Find(selector)
    if result.Length() == 0 {
        return "", fmt.Errorf("no elements found for selector: %s", selector)
    }

    html, err := result.Html()
    if err != nil {
        return "", err
    }
    result_trimmed := strings.TrimSpace(html)
    return result_trimmed, nil
}
