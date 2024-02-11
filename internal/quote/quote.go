package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Quote struct {
	QuoteUUID   string `json:"quote_uuid"`
	QuoteText   string `json:"quote_text"`
	QuoteAuthor string `json:"quote_author"`
}

type QuoteManager struct {
	Quotes []Quote
}

func NewQuoteManager() *QuoteManager {
	return &QuoteManager{}
}

func (qm *QuoteManager) LoadQuotesFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &qm.Quotes)
	if err != nil {
		return err
	}

	return nil
}

func (qm *QuoteManager) SaveQuotesToFile(filename string) error {
	data, err := json.MarshalIndent(qm.Quotes, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (qm *QuoteManager) AddQuote(quote Quote) {
	qm.Quotes = append(qm.Quotes, quote)
}

func (qm *QuoteManager) DeleteQuoteByUUID(uuid string) {
	for i, quote := range qm.Quotes {
		if quote.QuoteUUID == uuid {
			qm.Quotes = append(qm.Quotes[:i], qm.Quotes[i+1:]...)
			break
		}
	}
}

func (qm *QuoteManager) GetQuoteByUUID(uuid string) (Quote, error) {
	for _, quote := range qm.Quotes {
		if quote.QuoteUUID == uuid {
			return quote, nil
		}
	}
	return Quote{}, fmt.Errorf("quote not found")
}

func (qm *QuoteManager) ListAllQuotes() {
	for _, quote := range qm.Quotes {
		fmt.Printf("UUID: %s\n", quote.QuoteUUID)
		fmt.Printf("Text: %s\n", quote.QuoteText)
		fmt.Printf("Author: %s\n", quote.QuoteAuthor)
		fmt.Println("")
	}
}
