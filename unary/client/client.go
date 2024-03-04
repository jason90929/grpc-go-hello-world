package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/jason90929/grpc-go-hello-world/protos"
	"google.golang.org/grpc"
)

func main() {
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	c := protos.NewEchoServiceClient(cli)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("client started")

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		req := protos.EchoRequest{Req: string(line)}
		res, err := c.GetUnaryEcho(context.Background(), &req)
		if err != nil {
			panic(err)
		}
		fmt.Println("res:", res.GetRes())
	}
}
