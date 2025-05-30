package main

import (
	"log"
	"net/http"
	"quoteList/internal/quote"
)

func main() {

	router := http.NewServeMux()

	//handlers
	quote.NewQuoteHandler(router, quote.QuoteHandlerDeps{})
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Cant start server")
	}

}
