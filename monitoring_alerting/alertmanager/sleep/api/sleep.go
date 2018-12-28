package api

import (
	"net/http"
)

var isSleepingTime bool

// SleepHandler is a handler for set sleeping time.
func SleepHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		trigger, ok := r.URL.Query()["trigger"]

		isSleepingTime = false
		if ok && len(trigger) != 0 {
			isSleepingTime = false
			if trigger[0] == "1" {
				isSleepingTime = true
			}
		}

		rw.WriteHeader(http.StatusOK)
	}
}
