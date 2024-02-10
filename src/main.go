package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Quote struct {
	QuoteID     int    `json:"quote_id"`
	QuoteText   string `json:"quote_text"`
	QuoteAuthor string `json:"quote_author"`
}

type QuoteData struct {
	Data []Quote `json:"data"`
}

var quotes QuoteData

func main() {
	loadQuoteFile("./data/quote.json")

	http.HandleFunc("/add", addQuoteHandler)
	http.HandleFunc("/quote", getQuoteHandler)
	http.HandleFunc("/quotes", getAllQuotesHandler)

	fmt.Println("API server is running on :8090...")
	http.ListenAndServe(":8090", nil)
}

func loadQuoteFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &quotes)
}

func saveQuotesToFile(filename string) {
	data, err := json.Marshal(quotes)
	if err != nil {
		fmt.Println("Error marshalling quotes:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

func addQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newQuote Quote
		_ = json.NewDecoder(r.Body).Decode(&newQuote)

		// Add new quote
		quotes.Data = append(quotes.Data, newQuote)

		// save it
		saveQuotesToFile("./data/quote.json")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Param 'id' is missing", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(keys[0])
		if err != nil {
			http.Error(w, "Invalid id provided", http.StatusBadRequest)
			return
		}
		for _, quote := range quotes.Data {
			if quote.QuoteID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(quote)
				return
			}
		}
		http.Error(w, "Quote not found", http.StatusNotFound)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllQuotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
