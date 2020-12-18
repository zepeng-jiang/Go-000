package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	s := http.Server{Addr: "127.0.0.1:8080"}
	// http server
	g.Go(func() error {
		g.Go(func() error {
			<-ctx.Done()
			fmt.Println("http ctx done")
			return s.Shutdown(context.TODO())
		})
		fmt.Println("http start")
		return s.ListenAndServe()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal start")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return errors.New("signal exit")
			}
		}
	})

	err := g.Wait()
	fmt.Println("all", err)
}