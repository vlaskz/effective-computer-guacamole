package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type hwResponse struct {
	Message string `json:"message"`
}

type hwRequest struct {
	Name string `json:"name"`
}

const port = 8080

func main() {
	server()
}

func server() {

	http.HandleFunc("/hw", hwHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	log.Printf("Server started at %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func hwHandler(w http.ResponseWriter, r *http.Request) {

	var req hwRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	var msg = fmt.Sprintf("Hello %v", req.Name)
	res := hwResponse{Message: msg}

	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}
