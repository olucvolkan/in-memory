package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInMemoryHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/in-memory?key=abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	data := make(map[string]string)
	expected := "active-tabs"
	data["abc"] = expected
	handler := http.HandlerFunc(InMemoryGetHandler(&InMemoryKVStore{data: data}))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var inMemoryResponseHandler InMemoryResponse
	json.Unmarshal([]byte(rr.Body.String()), &inMemoryResponseHandler)
	if inMemoryResponseHandler.Value != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
