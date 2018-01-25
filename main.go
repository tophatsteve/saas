package main

import (
	"fmt"

	"github.com/tophatsteve/saas/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:9191",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := saas.NewSarcasmSvcClient(conn)
	ctx := context.Background()

	respGenerate, err := client.Generate(ctx, &saas.Empty{})

	fmt.Println(respGenerate.Value)
}
