package main

import (
	"cardValidator/pkg/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	addr             = "localhost:8080"
	writeTimeout     = 15 * time.Second
	readTimeout      = 15 * time.Second
	gracefulShutdown = 10 * time.Second
)

func main() {
	log.Println("Starting server...")

	var srv http.Server

	go func() {
		r := route.NewRoute()

		srv := &http.Server{
			Handler:      r,
			Addr:         addr,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		}

		log.Fatal(srv.ListenAndServe())
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdown)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Shutting down...")
	os.Exit(0)
}
