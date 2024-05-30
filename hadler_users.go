package main

import "net/http"

// make the handler a method to take in the apiConfig
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
