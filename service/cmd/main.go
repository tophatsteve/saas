package main

import (
	"net"
	"os"
	"os/signal"
    "google.golang.org/grpc"
    "github.com/tophatsteve/iaas/service"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":9191")

	if err != nil {
		panic(err)
	}
	
	grpcServer := grpc.NewServer()
	iaas.RegisterInsultSvcServer(grpcServer, iaas.NewService())

	go func() {
		grpcServer.Serve(lis)
	}()

	<-stop
	grpcServer.GracefulStop()
}