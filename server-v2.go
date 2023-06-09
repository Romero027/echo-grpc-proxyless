package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc/xds"

	echo "github.com/Romero027/echo-grpc-proxyless/pb"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, x *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("got: [%s]", x.GetReq())
	echoResponse := &echo.EchoResponse{
		Res:     x.GetReq(),
		Version: 1,
	}
	return echoResponse, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := xds.NewGRPCServer()

	fmt.Printf("Starting server at port 9000\n")

	echo.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
