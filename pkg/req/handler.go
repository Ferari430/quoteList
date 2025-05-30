package req

import (
	"net/http"
	"quoteList/internal/payload"
	"quoteList/pkg/res"
)

func HandleBody(w *http.ResponseWriter, r *http.Request) (*payload.Quote, error) {
	body, err := Decode(r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err

	}

	// VALIDATOR

	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
