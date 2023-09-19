package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jexodusmercado/blog-system-go/cmd"
	"github.com/jexodusmercado/blog-system-go/internal/api"
)

func main() {
	execCtx, execCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer execCancel()

	go func() {
		<-execCtx.Done()
		log.Println("received graceful shutdown signal")
	}()

	// command is expected to obey the cancellation signal on execCtx and
	// block while it is running
	if err := cmd.RootCommand().ExecuteContext(execCtx); err != nil {
		log.Fatalln(err)
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Minute)
	defer shutdownCancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("shutting down API servers")

		// wait for API servers to shut down gracefully
		api.WaitForCleanup(shutdownCtx)
	}()

	cleanupDone := make(chan struct{})
	go func() {
		defer close(cleanupDone)
		wg.Wait()
	}()

	select {
	case <-shutdownCtx.Done():
		// cleanup timed out
		return

	case <-cleanupDone:
		// cleanup finished before timing out
		return
	}

}