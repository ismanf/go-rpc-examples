package main;

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/ismayilmalik/go-rpc-examples/protofiles"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to listen on 8080: %v", err)
	}

	srv := grpc.NewServer()
	protofiles.RegisterAgeCheckerServer(srv, &Server{})
	log.Println("Starting server on port :8080")
	srv.Serve(listener)
}

type Server struct{}

func (s *Server) CheckAge(ctx context.Context, input *protofiles.CheckAgeRequest) (*protofiles.CheckAgeResponse, error)  {
	var result string

	if input.Data.Age > 35 {
		result = fmt.Sprintf("Hi %s, You are OLD!!!!!!", input.Data.Name)
	} else {
		result = fmt.Sprintf("Hi %s, You are YOUNG!!!!!!", input.Data.Name)
	}

	return &protofiles.CheckAgeResponse{ Result: result } , nil
}