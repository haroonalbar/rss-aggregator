package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// this function is used to handle the json response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// marshal means encoding.
  // this will convert the payload to json and return it as byte slice
  // which can be used to write on ResponseWriter
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON of %v",payload)
    // return response as internal server error
    w.WriteHeader(500)
		return
	}

  w.Write(data)
}