package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type WorkHandler struct{}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}

func (h *WorkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 5; i++ {
		log.Printf("Start %vth work.", i)
		time.Sleep(3 * time.Second)
		log.Printf("End %vth work.", i)
	}

	message := "Work done."
	if err := json.NewEncoder(w).Encode(message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
