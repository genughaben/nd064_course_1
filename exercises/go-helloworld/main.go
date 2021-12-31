package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type Response struct {
	Status string
}

func health(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: "OK - endpoint is up"}

	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func main() {
	http.HandleFunc("/status", health)
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":6111", nil)
}
