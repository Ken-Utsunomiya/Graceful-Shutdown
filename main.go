package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/projects/Graceful-Shutdown/handler/router"
)

func main() {
	addr := ":8080"
	wg := &sync.WaitGroup{}
	mux := router.NewRouter(wg)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen at %v: %s\n", addr, err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}

	log.Println("Waiting all processes finish...")
	wg.Wait()
	log.Println("All processes done.")

	log.Println("Server gracefully shutdown")
}
