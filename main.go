package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	SimpleWebServer()
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func SimpleWebServer() {
	http.HandleFunc("/ping", PingHandler)

	var message string = fmt.Sprintf("Server running on port %v", PORT)
	log.Print(message)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
