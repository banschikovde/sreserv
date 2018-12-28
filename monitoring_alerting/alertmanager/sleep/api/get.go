package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vlamug/elementary-cacher/cache"
	internalhttp "github.com/vlamug/elementary-cacher/pkg/http"
)

// GetHandler is a handler for loading data from cache.
func GetHandler(cacher cache.Cache) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		/// if it is sleeping time, sleep 1 second
		if isSleepingTime {
			time.Sleep(time.Second)
		}

		if r.Method != http.MethodPost {
			internalhttp.SendResponse(rw, GetResponse{Status: errorStatus, Message: "only POST requests are allowed"})
			return
		}

		req := GetRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		value, isExists := cacher.Get(req.Key)
		if !isExists {
			internalhttp.SendResponse(rw, GetResponse{Status: notFoundStatus, Message: "data by key do not found"})
			return
		}

		internalhttp.SendResponse(rw, GetResponse{Status: successStatus, Data: value})
	}
}

// GetRequest describes structure of request for loading data from cache.
type GetRequest struct {
	Key string `json:"key"`
}

// GetRequest describes structure of response with loaded data from cache.
type GetResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
