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
	connection, err := grpc.Dial(apiHost+":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot connect to server")
	}
	client = amassa.NewInputHandlerServiceClient(connection)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/work", mainHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
