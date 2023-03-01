package router

import (
	"net/http"
	"sync"

	"github.com/projects/Graceful-Shutdown/handler"
	"github.com/projects/Graceful-Shutdown/handler/middleware"
)

func NewRouter(wg *sync.WaitGroup) *http.ServeMux {
	mux := http.NewServeMux()
	h := handler.NewWorkHandler()
	mux.Handle("/", middleware.TrackProcess(h, wg))
	return mux
}
