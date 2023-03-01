package middleware

import (
	"net/http"
	"sync"
)

func TrackRequest(h http.Handler, wg *sync.WaitGroup) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		defer wg.Done()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
