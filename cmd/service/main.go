package main

import (
	"flag"
	"fmt"
	"gRPC_study/pb"
	"gRPC_study/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()
	log.Printf("start server on port: %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("listen address: %v", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start listener: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server: ", err)
	}

}
