package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server Starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	response := helloWorldResponse{Message: "You saluted the world, " + request.Name, Id: i}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	i++
}
