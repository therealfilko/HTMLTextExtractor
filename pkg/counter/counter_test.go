package counter

import (
    "testing"
)

func TestCountWords(t *testing.T) {
    count, err := CountWords("http://example.com")
    if err != nil {
        t.Fatalf("Fehler: %v", err)
    }
    if count == 0 {
        t.Errorf("Erwartete Wörter zu zählen, aber erhielt 0")
    }
}
