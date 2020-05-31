package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"tweeter/router"
)

func main() {
	handler := router.Handle()

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

    router.ConnectToDB("mongodb://app:27017")

	http2.ConfigureServer(&server, &http2.Server{})
	log.Fatal(server.ListenAndServe())
}
