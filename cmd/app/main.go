package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_server "github.com/2yanpath/grpc-error-detail-test/internal/presentation/grpc"
	greetpb "github.com/2yanpath/grpc-error-detail-test/proto/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8081
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, grpc_server.NewGreetServer())
	reflection.Register(s)
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
