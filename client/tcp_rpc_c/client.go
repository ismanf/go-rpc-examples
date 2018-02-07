package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type GreetArgs struct {
	FirstName, LastName string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:3001")
	if err != nil {
		log.Fatal("dial error...", err)
	}

	args := GreetArgs{"John", "Doe"}
	var message string

	err = client.Call("Greeter.Greet", args, &message)
	if err != nil {
		log.Fatal("call error...", err)
	}

	fmt.Println("response:", message)
}
