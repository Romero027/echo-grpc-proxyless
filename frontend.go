package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/xds"

	echo "github.com/Romero027/echo-grpc-proxyless/pb"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s\n", request.URL.String())

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("xds:///echo-server:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := echo.NewEchoServiceClient(conn)

	message := echo.EchoRequest{
		Req: request.URL.String(),
	}

	response, err := c.Echo(context.Background(), &message)
	if err != nil {
		log.Fatalf("Erro when calling echo: %s", err)
	}
	if err == nil {
		log.Printf("Response from server: %s", response.Res)
		fmt.Fprintf(writer, "Echo request finished! Length of the request is %d\n", len(response.Res))
	} else {
		log.Printf("Erro when calling echo: %s", err)
		fmt.Fprintf(writer, "Echo server returns an error: %s\n", err)
	}
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Printf("Starting frontend at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
