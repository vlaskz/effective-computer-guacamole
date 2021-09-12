package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var port = 8080
var i = 0

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {

	rnd := renderer.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := helloWorldResponse{Message: "Hello World", Id: i}
		rnd.JSON(w, http.StatusOK, res)
		i++
	})

	mux.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server Starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), mux))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "You saluted the world, " + request.Name, Id: i}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	i++
}
