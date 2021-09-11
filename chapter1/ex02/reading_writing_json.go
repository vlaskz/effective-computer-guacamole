package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port = 8080
var i = 0

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	Date    string `json:", omitempty"`
	Id      int    `json:"id, string"`
}

func main() {

	http.HandleFunc("/helloworld", helloWorldHandled)

	log.Printf("Server Starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandled(w http.ResponseWriter, r *http.Request) {

	response := helloWorldResponse{Message: "helloWorld Still the best!", Id: i}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	i++
}
