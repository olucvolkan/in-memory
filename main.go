package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func main() {
	config := newConfig()

	kvstore := NewInMemoryKVStore()
	http.HandleFunc("/in-memory", buildInMemoryHandler(kvstore))

	log.Println("Starting Server")
	e := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	log.Fatal(e)
}

func newConfig() *Config {
	config := &Config{}

	if os.Getenv("PORT") == "" {
		log.Fatal("Wrong PORT info, failed to start the app")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Can't parse PORT, failed to start the app")
	}
	config.Port = port

	return config
}
