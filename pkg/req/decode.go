package req

import (
	"encoding/json"
	"io"
	"quoteList/internal/payload"
)

func Decode(body io.ReadCloser) (payload.Quote, error) {
	var payload payload.Quote 

	err := json.NewDecoder(body).Decode(&payload) 
	
	if err != nil {
		return payload, err
	}
	return payload, nil
}
