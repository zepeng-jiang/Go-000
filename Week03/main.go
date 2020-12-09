package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var stop = make(chan string)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	// 启动两个服务
	group.Go(userServer)
	group.Go(orderServer)
	// signal信号注册
	signalRegister()

	<-ctx.Done()
	if err := ctx.Err(); err != nil {
		log.Println("select received: ", err)
		if stop != nil {
			stop <- "stop"
		}
	}
	close(stop)
	time.Sleep(time.Second * 2)
	log.Println("all server has shutdown")
}

// userServer 用户服务
func userServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user server received successfully"))
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	go func() {
		<-stop
		log.Println("user server received close signal")
		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(timeout)
		log.Println("user server has shutdown")
	}()

	log.Println("user server: listening on port :8081")
	return server.ListenAndServe()
}

// orderServer 订单服务
func orderServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/order", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("order server received successfully"))
	})

	server := http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	go func() {
		<-stop
		log.Println("order server received close signal")
		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(timeout)
		log.Println("order server has shutdown")
	}()

	log.Println("order server: listening on port :8082")
	return server.ListenAndServe()
}

// signalRegister 注册signal信号
func signalRegister() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()

		si := <-signals
		log.Println("Got signal :", si)
		if stop != nil {
			stop <- "stop"
		}
	}()
}
