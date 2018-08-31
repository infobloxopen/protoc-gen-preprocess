package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	demo "github.com/infobloxopen/protoc-gen-preprocess/example/proto"
	grpc_preprocessor "github.com/infobloxopen/protoc-gen-preprocess/middleware"
	"google.golang.org/grpc"
)

const (
	demoAddress   = "localhost:8081"
	serverAddress = "localhost:8080"
)

type demoServer struct{}

//Echo implementation of demo service
func (s demoServer) Echo(ctx context.Context, d *demo.Demo) (*demo.Demo, error) {
	return &demo.Demo{PreprocessedField: d.GetPreprocessedField(), Untouched: d.GetUntouched()}, nil
}

func runService() {
	// Middleware chain.
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_preprocessor.UnaryServerInterceptor(), // preprocessing middleware
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
	demo.RegisterDemoServiceServer(server, &demoServer{})
	ln, err := net.Listen("tcp", demoAddress)
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(ln); err != nil {
		log.Fatal(err)
	}
}

func newDemoHandler(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := demo.RegisterDemoServiceHandlerFromEndpoint(ctx, mux, addr, dialOpts)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func runServer() {
	mux := http.NewServeMux()
	errHandler := runtime.WithProtoErrorHandler(gateway.ProtoMessageErrorHandler)
	opHandler := runtime.WithMetadata(gateway.MetadataAnnotator)

	demoHandler, err := newDemoHandler(context.Background(), demoAddress, errHandler, opHandler)
	if err != nil {
		log.Fatalln(err)
	}
	mux.Handle("/", demoHandler)
	http.ListenAndServe(serverAddress, mux)
}

func main() {
	go runService()
	runServer()
}
