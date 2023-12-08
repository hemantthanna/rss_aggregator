package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/// Note: error codes below 499 are client side error. thats their problem
/// 	  but codes above 500 level are our's to handle.



func respondWithError(w http.ResponseWriter, code int, message string)  {
	if code > 499 {
		log.Println("Responding with 5XX error:", message)
	}
	
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: message,
	})
}


// function to respond with a jsondata
func respondWithJSON(w http.ResponseWriter, code int, payload interface{})  {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) // internal server error
		return
		
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) // everything went well
	w.Write(data)
	
}