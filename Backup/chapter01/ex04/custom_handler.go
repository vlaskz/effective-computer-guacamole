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

type validationHandler struct {
	next http.Handler
}

type hwHandler struct{}

func main() {
	port := 8080

	handler := newValidationHandler(newHWHandler())

	http.Handle("/hw", handler)

	log.Printf("Serving at port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var req hwRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(rw, r)
}

func newHWHandler() http.Handler {
	return hwHandler{}
}

func (h hwHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	res := hwResponse{Message: "Hello!"}
	fmt.Println("log!")

	encoder := json.NewEncoder(rw)
	encoder.Encode(res)
}
