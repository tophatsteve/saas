package main

import (
	"fmt"

	saas "github.com/tophatsteve/saas/service"
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

	client := saas.NewSausageSvcClient(conn)
	ctx := context.Background()

	respGenerate, err := client.Generate(ctx, &saas.Empty{})

	if err != nil {
		panic(err)
	}

	fmt.Println(respGenerate.Value)
}
