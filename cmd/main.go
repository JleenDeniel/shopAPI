package main

import (
	"context"
	pb "example.com/mymodule/gen/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type shopAPiServer struct {
	pb.UnimplementedShopAPIServer
}

func (s *shopAPiServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	return &pb.AddToCartResponse{}, nil
}

func main() {
	go func() {
		mux := runtime.NewServeMux()
		err := pb.RegisterShopAPIHandlerServer(context.Background(), mux, &shopAPiServer{})
		if err != nil {
			return
		}
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShopAPIServer(grpcServer, &shopAPiServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
