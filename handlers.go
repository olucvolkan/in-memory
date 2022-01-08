package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

const (
	SuccessStatus int = 0
	FailStatus    int = 1
)

type InMemoryRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InMemoryResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

type InMemoryMap struct {
	KeyValuePair map[string]string
	Mutex        *sync.Mutex
}

var customMap *InMemoryMap

func buildInMemoryHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			InMemoryPostHandler(kvstore)(w, r)
		} else if r.Method == "GET" {
			InMemoryGetHandler(kvstore)(w, r)
		}
	}
}

func InMemoryPostHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request InMemoryRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&InMemoryResponse{Code: FailStatus, Message: "invalid request"})
			return
		}

		if ok, err := kvstore.Set(request.Key, request.Value); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&InMemoryResponse{Code: FailStatus, Message: err.Error()})
			return
		}

		// Encode results
		json.NewEncoder(w).Encode(&InMemoryResponse{
			Key:   request.Key,
			Value: request.Value,
		})
	}
}

func InMemoryGetHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code": FailStatus,
				"msg":  "Invalid request",
			})
			return
		}

		key := r.URL.Query().Get("key")
		value, err := kvstore.Get(key)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&InMemoryResponse{Code: FailStatus, Message: err.Error()})
			return
		}
		// Encode results
		json.NewEncoder(w).Encode(&InMemoryResponse{
			Key:   key,
			Value: value,
		})
	}
}
