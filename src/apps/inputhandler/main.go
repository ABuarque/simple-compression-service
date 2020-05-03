package main

import (
	"context"
	"encoding/json"
	"fmt"
	amassa "inputhandler/proto"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ABuarque/simple-compression-service/src/libs/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// it is the message that will be published
// to be compressed or decompressed
type in struct {
	File    string `json:"file"`
	Email   string `json:"email"`
	Command string `json:"command"`
}

var re *redis.Client

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in in
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		log.Println("error decoing payload from request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	payloadAsBytes, err := json.Marshal(in)
	if err != nil {
		log.Println("error marshaling payload to pusblish: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	payload := make(map[string]string)
	if in.Command == "compress" {
		re.Publish("compression", string(payloadAsBytes))
		log.Println("published compression topic")
		payload["msg"] = "file being compressed!"
	} else {
		re.Publish("decompression", string(payloadAsBytes))
		log.Println("published decompression topic")
		payload["msg"] = "file being decompressed!"
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

type Server struct{}

// RequestAction is a
func (s *Server) RequestAction(ctx context.Context, req *amassa.Request) (*amassa.Response, error) {
	return &amassa.Response{}, nil
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	re = redis.New()
	// http.HandleFunc("/process", handleRequest)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	grpcServer := grpc.NewServer()
	amassa.RegisterInputHandlerServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
