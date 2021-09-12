package main

import (
	"fmt"
	"log"
	"net/http"
)

var i int = 0
var port int = 8080

func main() {

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World #%v\n", i)
	log.Printf("This was the request #%v", i)
	i++
}
