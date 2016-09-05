package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type message struct {
	Payload string `json:payload`
}

func parsedPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	
	var t message
	err := decoder.Decode(&t)
	
	if err != nil {
		panic(err)
	}	

	fmt.Println("result: ", t.Payload)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/events/", parsedPost)
	r.Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", r))
}
