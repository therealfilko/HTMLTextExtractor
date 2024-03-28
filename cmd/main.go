package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Was soll das programm können?

// Man soll eine url und ein css selektor angeben | DONE
// Als Ergebnis soll dann der Inhalt des Tags sein
// z.B.: extractor https://pixelding.de h1
// Output soll "Die Agentur, die du suchst."

// os.Args ist die Methode um mit Argumente im Aufruf des Programmes zu speichern und zu verwenden

func main() {
    url := os.Args[1]
    selektor := os.Args[2]

    fmt.Println(url)
    fmt.Println(selektor)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
    }

    //Load the HTML Document
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
}

// Wir können Argumente verwenden und speichern. Dafür haben wir jetzt die Variable url und selektor
// Wir können jetzt eine Anfrage machen
// Anfragen klappen und body wird gespeichert, aber da müssen wir noch mit goquery arbeiten und testen
