package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInMemoryHandler(t *testing.T) {
	//Given
	req, err := http.NewRequest("GET", "/in-memory?key=abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	data := make(map[string]string)
	expected := "active-tabs"
	data["abc"] = expected
	//When
	handler := http.HandlerFunc(InMemoryGetHandler(&InMemoryKVStore{data: data}))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//Then
	var inMemoryResponseHandler InMemoryResponse
	json.Unmarshal([]byte(rr.Body.String()), &inMemoryResponseHandler)
	if inMemoryResponseHandler.Value != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInMemoryKVStore_FlushAll(t *testing.T) {
	//Given
	var jsonData = []byte(`{
		"key": "test",
		"value": "value"
	}`)
	req, err := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	data := make(map[string]string)
	//When
	handler := http.HandlerFunc(InMemoryFlushAllHandler(&InMemoryKVStore{data: data}))
	handler.ServeHTTP(rr, req)
	//Then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var inMemoryResponseHandler InMemoryResponse
	json.Unmarshal([]byte(rr.Body.String()), &inMemoryResponseHandler)
	if inMemoryResponseHandler.Code != SuccessStatus {
		t.Errorf("handler returned unexpected body: got %v ",
			rr.Body.String())
	}
}

func TestInMemoryPostHandler(t *testing.T) {
	//Given
	var jsonData = []byte(`{
		"key": "test",
		"value": "value"
	}`)
	req, err := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	data := make(map[string]string)
	expected := "value"
	data["test"] = expected
	//When
	handler := http.HandlerFunc(InMemoryPostHandler(&InMemoryKVStore{data: data}))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Then
	var inMemoryResponseHandler InMemoryResponse
	json.Unmarshal([]byte(rr.Body.String()), &inMemoryResponseHandler)
	if inMemoryResponseHandler.Value != expected {
		t.Errorf("handler returned unexpected body: got %v  ",
			rr.Body.String())
	}
}
