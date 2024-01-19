package main

import (
	"broker/internal/routes"
	"fmt"
	"log"
	"net/http"
)

const WebPort = "80"

type Config struct {
	Routes http.Handler
}

func main() {
	app := Config{
		Routes: routes.Routes(),
	}

	log.Printf("Starting Broker Services at port : %s\n", WebPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.Routes,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
