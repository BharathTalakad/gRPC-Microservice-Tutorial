package main

import (
	"fmt"
	"github.com/bharath/go-grpc-http-rest-microservice-tutorial/servers"
	"os"
)

func main() {
	if err := cmd.RunServer() ; err != nil {
		fmt.Fprintf(os.Stderr, "%v\n",err)
		os.Exit(1)
	}
}
