package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) { // data interface{} - it can take any type of data (int, float, struct, list etc).
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, "Your response: ")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}