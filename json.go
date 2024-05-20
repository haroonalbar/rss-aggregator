package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errResponse struct {
		// `json:"error"` tag here is used to tell how to marshal or unmarshal this struct
		// here it will be attached to key "error" : value what ever the Error is
		Error string `json:"error"`
	}
	// The json.Marshal will look something like this.
	// {
	//   "error": "something something"
	// }
	respondWithJSON(w, code, errResponse{Error: msg})
}

// this function is used to handle the json response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// marshal means encoding.
	// this will convert the payload to json and return it as byte slice
	// which can be used to write on ResponseWriter
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON of %v", payload)
		// return response as internal server error
		w.WriteHeader(500)
		return
	}

	// added a response header
	// saying we are responding in json
	w.Header().Add("Content-Type", "application/json")
	// added the status quote on response header
	w.WriteHeader(code)
	// finally write the data on response body
	w.Write(data)
}
