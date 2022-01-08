package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	logPath := "development.log"
	kvstore := NewInMemoryKVStore()
	kvstore = addFileData(kvstore)

	openLogFile(logPath)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/in-memory", buildInMemoryHandler(kvstore))

	log.Println("Starting Server")
	e := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), logRequest(http.DefaultServeMux))
	log.Fatal(e)
}

// addFileData append json file data in store
func addFileData(store KVStore) KVStore {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		return store
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	for key, value := range result {
		store.Set(key, value.(string))
	}
	return store
}

// newConfig create web server config
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

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n\n", r.RemoteAddr, r.Method, r.URL)
		log.Printf("%s %s %s\n\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}
