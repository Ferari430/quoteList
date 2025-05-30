package res_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"quoteList/pkg/res"
)

func TestJson(t *testing.T) {

	rr := httptest.NewRecorder()

	data := map[string]string{"foo": "bar"}

	status := http.StatusCreated // 201

	err := res.Json(rr, data, status)
	if err != nil {
		t.Fatalf("Json returned error: %v", err)
	}

	if rr.Code != status {
		t.Errorf("Expected status %d, got %d", status, rr.Code)
	}

	ct := rr.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}

	expectedBody := `{"foo":"bar"}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
	}
}
