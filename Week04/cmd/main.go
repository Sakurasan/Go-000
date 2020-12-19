package main

import (
	"Go-000/Week04/handlers"
	"Go-000/Week04/pkg/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.Server{Addr: config.GetString("App.addr"), Handler: handlers.NewSvr()}

	// http server
	g.Go(func() error {
		fmt.Println("http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			svr.Shutdown(context.TODO())
		}()
		if err := svr.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Fatal("Server closed under request")
			} else {
				log.Fatal("Server closed unexpect")
			}
		}
		return nil
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				if err := svr.Shutdown(context.TODO()); err != nil {
					log.Fatal("Server forced to shutdown:", err)
				}
				handlers.Close()
				return nil
			}
		}
	})

	// inject error
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}
