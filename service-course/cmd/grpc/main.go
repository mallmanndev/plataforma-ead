package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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
	// pb.RegisterUsersServiceServer(grpcServer, &grpcadapter.UsersServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
