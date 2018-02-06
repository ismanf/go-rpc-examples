package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type GreetArgs struct {
	FirstName, LastName string
}

type Greeter struct{}

func (g *Greeter) Greet(args *GreetArgs, message *string) error {
	*message = fmt.Sprintf("Hello %s %s! Greetings..", args.FirstName, args.LastName)
	return nil
}

func main() {
	greeter := new(Greeter)
	rpc.Register(greeter)
	rpc.HandleHTTP()

	fmt.Println("HTTP-RPC service started on port:", "3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
