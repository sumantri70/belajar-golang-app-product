package helper

import (
	"encoding/json"
	"net/http"
)

func WriteResponseJson(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	PanicIfError(encoder.Encode(response))
}
