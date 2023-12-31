package main

import (
	"fmt"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/adapters/grpc"
	"log"
	"net"

	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC server!!!")

	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("Error to listen server in port")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, &grpcroutes.UsersServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
