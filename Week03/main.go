package mian

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {

	server := &http.Server{
		Addr:    "0.0.0.0:" + "8080",
		Handler: &router{},
	}

	handleSignal(server)

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("listen and serve failed: ", err.Error())
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); nil != err {
			logrus.Errorf("server close failed: " + err.Error())
		}

		// DB.DisconnectDB()
		os.Exit(0)
	}()
}

type router struct{}

func (*router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("1")
		return nil
	})

	g.Go(func() error {
		fmt.Println("2")
		return nil
	})

	g.Go(func() error {
		fmt.Println("3")
		return nil
	})
	if err := g.Wait(); err != nil {
		logrus.Errorln(err)
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	w.Write([]byte("Hello, world!"))
}
