package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestJson(httpRequest *http.Request, request any) {
	decoder := json.NewDecoder(httpRequest.Body)
	PanicIfError(decoder.Decode(request))

}
