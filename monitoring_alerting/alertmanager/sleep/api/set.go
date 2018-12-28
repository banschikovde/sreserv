package api

import (
	"encoding/json"
	"net/http"

	"github.com/vlamug/elementary-cacher/cache"
	internalhttp "github.com/vlamug/elementary-cacher/pkg/http"
)

// SetHandler is a handler for setting data into cache
func SetHandler(cacher cache.Cache) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			internalhttp.SendResponse(rw, SetResponse{Status: errorStatus, Message: "only POST requests are allowed"})
			return
		}

		req := SetRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(req.Key) == 0 {
			internalhttp.SendResponse(rw, SetResponse{Status: errorStatus, Message: "incorrect key value"})
			return
		}

		cacher.Set(req.Key, req.Value)

		internalhttp.SendResponse(rw, SetResponse{Status: successStatus})
	}
}

// SetRequest describes structure of request for setting data into cache.
type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetResponse describes structure of response of setting data into cache.
type SetResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
