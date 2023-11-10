package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	grpc2 "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"google.golang.org/grpc"
)

func main() {
	var port = 3000
	fmt.Println("Starting gRPC server!")

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error to listen server in port %s", strconv.Itoa(port))
	}

	db, disconnect := utils.GetDb()
	defer disconnect()

	// REGISTER ROUTES HERE
	grpcServer := grpc.NewServer()
	coursesServer := grpc2.NewCourseServer(db)
	fileUploadService := grpc2.NewFilesServer(db)
	pb.RegisterCoursesServiceServer(grpcServer, coursesServer)
	pb.RegisterFileUploadServiceServer(grpcServer, fileUploadService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
