package main

import (
	"encoding/json"
	"net/http"
)

const (
	SuccessStatus int = 0
	FailStatus    int = 1
)

type InMemoryPostRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type InMemoryGetRequest struct {
	Key string `json:"key"`
}

type InMemoryResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

// handle request
func buildInMemoryHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			InMemoryPostHandler(kvstore)(w, r)
		} else if r.Method == "GET" {
			InMemoryGetHandler(kvstore)(w, r)
		} else if r.Method == "DELETE" {
			InMemoryFlushAllHandler(kvstore)(w, r)
		}
	}
}

// InMemoryPostHandler create key value endpoint
// PostInMemory godoc
// @Summary create in memory record
// @Description create in memory record
// @Tags inMemory
// @Param tasks body InMemoryPostRequest true "in memory info"
// @Accept  json
// @Produce  json
// @Router /in-memory [post]
func InMemoryPostHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request InMemoryPostRequest
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

// InMemoryFlushAllHandler flush all in-memory values
// DeleteInMemory godoc
// @Summary delete in memory records
// @Description delete in memory records
// @Tags inMemory
// @Accept  json
// @Produce  json
// @Router /in-memory [delete]
func InMemoryFlushAllHandler(kvstore KVStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := kvstore.FlushAll()
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code": FailStatus,
				"msg":  "Failed flush data ",
			})
		}
		// Encode results
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": SuccessStatus,
			"msg":  "Removed All Data",
		})
	}
}

// InMemoryGetHandler get  key value endpoint
// GetInMemory godoc
// @Summary list in memory records
// @Description list in memory records
// @Tags inMemory
// @Accept  json
// @Param tasks query InMemoryGetRequest true "in memory info"
// @Produce  json
// @Router /in-memory [get]
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
