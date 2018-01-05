package main

import (
	"fmt" 
	"io"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "github.com/tophatsteve/iaas/service"   
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	//"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"	
)

func main() {
    tracer, closer := initTracer("iaas-client")
    opentracing.SetGlobalTracer(tracer)
    defer closer.Close()

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
    
    span := tracer.StartSpan("generate")
    defer span.Finish()
    respGenerate, err := client.Generate(ctx, &iaas.Empty{})
    span.SetTag("response", respGenerate.Value)

    fmt.Println(respGenerate.Value)
}

func initTracer(service string) (opentracing.Tracer, io.Closer) {
    cfg := &config.Configuration{
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
        },
    }
    tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    return tracer, closer
}