package quote

import (
	"encoding/json"
	"log"
	"net/http"
	"quoteList/internal/payload"
	"quoteList/pkg/req"
	"quoteList/pkg/res"
	"quoteList/repo"
	"strconv"
)

type QuoteHandler struct {
	storage repo.Store
}

type QuoteHandlerDeps struct {
	storage repo.Store
}

func NewQuoteHandler(router *http.ServeMux, deps QuoteHandlerDeps) {
	handler := QuoteHandler{}

	router.HandleFunc("POST /quotes", handler.Create())
	router.HandleFunc("GET  /quotes", handler.GetQuotes())
	router.HandleFunc("GET  /quotes/random", handler.GetRandom())
	router.HandleFunc("DELETE /quotes/{id}", handler.Delete())

}

func (handler *QuoteHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody(&w, r)
		if err != nil {
			log.Printf("Cant decode body, %v\n", err.Error())
			res.Json(w, "Cant decode body", http.StatusBadRequest)
			return
		}

		res.Json(w, body, http.StatusCreated)

		handler.storage.Create(body)
	}
}

func (handler *QuoteHandler) GetQuotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")

		var quotes []payload.Quote
		if author == "" {
			quotes = handler.storage.GetAll()
			if len(quotes) == 0 {
				res.Json(w, "No quotes found", http.StatusNotFound)
				return
			}

		} else {
			quotes = handler.storage.Filter(author)
			if len(quotes) == 0 {
				res.Json(w, "No quotes found with this author", http.StatusNotFound)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(quotes); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func (handler *QuoteHandler) GetRandom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		randomQuote, err := handler.storage.GetRandom()

		if err != nil {
			res.Json(w, "No quotes found", http.StatusNotFound)
			return
		}

		res.Json(w, randomQuote, http.StatusOK)
	}
}

func (handler *QuoteHandler) Filter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		quotes := handler.storage.Filter(author)
		res.Json(w, quotes, http.StatusOK)
	}
}

func (handler *QuoteHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Cant parse id to uint", http.StatusBadGateway)
			return
		}

		handler.storage.Delete(id)
		res.Json(w, "Quote deleted", http.StatusOK)
	}
}
