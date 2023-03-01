package middleware

import (
	"net/http"
	"sync"
)

func TrackProcess(h http.Handler, wg *sync.WaitGroup) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// increment when a new request comes in, dones when response sent
		wg.Add(1)
		defer wg.Done()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
