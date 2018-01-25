package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/tophatsteve/saas/service"
	"google.golang.org/grpc"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":9191")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	saas.RegisterSarcasmSvcServer(grpcServer,
		saas.NewService(
			[]string{"Your mother was a hamster, and your father smelt of elderberries."},
		),
	)

	go func() {
		grpcServer.Serve(lis)
	}()

	<-stop
	grpcServer.GracefulStop()
}
