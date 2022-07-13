package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, requestBody interface{}) {
	// decode data dari request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(requestBody)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, responseBody interface{}) {
	// set header
	w.Header().Add("Content-Type", "application-json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseBody)
	PanicIfError(err)
}
