package main

import (
	"fmt"
	grpcAdapter "github.com/matheusvmallmann/plataforma-ead/service-core/aplication/adapters/grpc"
	"log"
	"net"
	"strconv"

	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
	"google.golang.org/grpc"
)

func main() {
	var port = 3000
	fmt.Println("Starting gRPC server!")

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error to listen server in port %s", strconv.Itoa(port))
	}

	// REGISTER ROUTES HERE
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, &grpcAdapter.UsersServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
