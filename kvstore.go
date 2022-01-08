package main

import "errors"

type KVStore interface {
	Get(key string) (string, error)
	Set(key string, value string) (bool, error)
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
	return true, nil
}

func NewInMemoryKVStore() KVStore {
	return &InMemoryKVStore{data: make(map[string]string)}
}
