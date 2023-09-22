package main

import (
	"fmt"
	grpc2 "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
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

	db, disconnect := utils.GetDb("dev")
	defer disconnect()

	// REGISTER ROUTES HERE
	grpcServer := grpc.NewServer()
	coursesServer := grpc2.NewCourseServer(db)
	pb.RegisterCoursesServiceServer(grpcServer, coursesServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
