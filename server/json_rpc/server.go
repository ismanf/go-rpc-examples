package main

import (
	"net/rpc/jsonrpc"
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

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3002")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
