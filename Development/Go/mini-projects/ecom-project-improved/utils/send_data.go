package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendData sets the HTTP status code and serializes data as JSON to the response writer.
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, "Your response: ")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}