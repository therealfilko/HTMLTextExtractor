package main

import (
    "fmt"
    "log"
    "os"

    "github.com/therealfilko/HTMLTextExtractor/pkg/extractor"
)

func main() {
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
