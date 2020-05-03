package main

import (
	"log"
	"net/http"
	"os"

	amassa "frontend/proto"

	"google.golang.org/grpc"
)

// holds the api ip
var apiHost string

var client amassa.InputHandlerServiceClient

func main() {
	apiHost = os.Getenv("API_HOST")
	if apiHost == "" {
		apiHost = "localhost"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8086"
	}
	log.Println("host: ", apiHost)
	connection, err := grpc.Dial(apiHost+":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot connect to server")
	}
	client = amassa.NewInputHandlerServiceClient(connection)
	log.Println("created new gRPC client")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/work", mainHandler)
	log.Println("server online at ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
