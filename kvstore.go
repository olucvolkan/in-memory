package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type KVStore interface {
	Get(key string) (string, error)
	Set(key string, value string) (bool, error)
	FlushAll() (bool, error)
}

type InMemoryKVStore struct {
	data map[string]string
}

func (s InMemoryKVStore) Get(key string) (string, error) {
	value, ok := s.data[key]
	if !ok {
		return "", errors.New("key not found")
	}

	return value, nil
}

func (s InMemoryKVStore) Set(key string, value string) (bool, error) {
	s.data[key] = value
	go writeFile(s.data)
	return true, nil
}
func writeFile(data map[string]string) {
	file, _ := json.MarshalIndent(data, "", " ")
	fmt.Println("test")
	ioutil.WriteFile("test.json", file, 0644)
}

func (s InMemoryKVStore) FlushAll() (bool, error) {
	for k := range s.data {
		delete(s.data, k)
	}
	file, _ := json.MarshalIndent(s.data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	return true, nil
}

func NewInMemoryKVStore() KVStore {
	return &InMemoryKVStore{data: make(map[string]string)}
}
