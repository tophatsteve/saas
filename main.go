package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "github.com/tophatsteve/iaas/service"   
    "fmt" 
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
    
	client := iaas.NewInsultSvcClient(conn)
	ctx := context.Background()
    respGenerate, err := client.Generate(ctx, &iaas.Empty{})
    
    fmt.Println(respGenerate.Value)
}
