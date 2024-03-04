package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jason90929/grpc-go-hello-world/protos"
	"google.golang.org/grpc"
)

type echoService struct {
	protos.UnimplementedEchoServiceServer
}

func (es *echoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	r := req.GetReq()
	fmt.Println("received:", r)
	s := "received-" + r
	res := protos.EchoResponse{
		Res: s,
	}
	return &res, nil
}

func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, new(echoService))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	rpcs.Serve(lis)

	fmt.Println("server started at 127.0.0.1:8080")
}
