package main

import (
    "fmt"
    "log"
    "os"

    extractor "github.com/therealfilko/HTMLTextExtractor/pkg"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatal("Usage: go run main.go <URL> <selector>")
    }

    url := os.Args[1]
    selector := os.Args[2]

    fmt.Println(url)
    fmt.Println(selector)

    result, err := extractor.FetchAndExtractHTML(url, selector)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
}
