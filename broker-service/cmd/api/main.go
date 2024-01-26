package main

import (
	broker_controller "broker/internal/controller/broker"
	http_server "broker/internal/http"
	"broker/internal/http/middleware"
	"fmt"
	"log"
)

const WebPort = "80"

func main() {
	log.Printf("Starting Broker Services at port : %s\n", WebPort)

	srv := http_server.NewHTTPServer("debug")

	root := srv.Group("")

	root.Use(middleware.ErrorHandler())

	broker_controller.BrokerController(root)

	srv.Run(fmt.Sprintf(":%s", WebPort))
}
