package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
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

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3001")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	rpc.Accept(listener)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
