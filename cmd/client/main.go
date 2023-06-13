package main

import (
	"context"
	"flag"
	"gRPC_study/pb"
	"gRPC_study/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	serverAddress := flag.String("address", "127.0.0.1:8080", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{Laptop: laptop}

	resp, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		statusCode, ok := status.FromError(err)
		if ok && statusCode.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Print("created laptop with id: ", resp.Id)
}
