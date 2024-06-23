package main

import (
    "fmt"
    "log"
    "os"

    extractor "github.com/therealfilko/HTMLTextExtractor/pkg/extractor"
    counter "github.com/therealfilko/HTMLTextExtractor/pkg/counter"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatal("Usage: go run main.go <mode> <URL> [<selector>]")
    }

    mode := os.Args[1]
    url := os.Args[2]

    switch mode {
    case "extract":
        if len(os.Args) < 4 {
            log.Fatal("Usage: go run main.go extract <URL> <selector>")
        }
        selector := os.Args[3]
        result, err := extractor.FetchAndExtractHTML(url, selector)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(result)
    case "count":
        result, err := counter.CountWords(url)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Word count:", result)
    default:
        log.Fatal("Unknown mode:", mode)
    }
}
