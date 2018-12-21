package http

import (
	"encoding/json"
	"net/http"
)

func SendResponse(rw http.ResponseWriter, body interface{}) {
	rw.Header().Add("Content-Type", "application/json")

	resp, err := json.Marshal(body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(resp)
}
