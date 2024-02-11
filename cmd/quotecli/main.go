package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ac227/go-catquote/internal/quote"
)

func main() {
	qm := quote.NewQuoteManager()

	// Load quotes from file
	err := qm.LoadQuotesFromFile("data/data.json")
	if err != nil {
		log.Fatal("Error loading quotes:", err)
	}
	uuid, _ := generateUUID()

	// Example usage
	quote1 := quote.Quote{
		QuoteUUID:   uuid,
		QuoteText:   "A",
		QuoteAuthor: "Sd",
	}

	// Add quote
	qm.AddQuote(quote1)

	// List all quotes
	qm.ListAllQuotes()

	// Save quotes to file
	err = qm.SaveQuotesToFile("data/data.json")
	if err != nil {
		log.Fatal("Error saving quotes:", err)
	}
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	// Set version (4) and variant (2)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
