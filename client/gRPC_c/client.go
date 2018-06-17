package main;

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/ismayilmalik/go-rpc-examples/protofiles"
	"log"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	client := protofiles.NewAgeCheckerClient(conn)

	userData := &protofiles.UserData{ Age: 40, Name: "Ehmed" }
	checkRequest := &protofiles.CheckAgeRequest{ Data: userData }
	resp, err := client.CheckAge(context.Background(), checkRequest)
	fmt.Printf("Response: %s \n", resp.Result)
}