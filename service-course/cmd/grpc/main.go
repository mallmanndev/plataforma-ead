package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	grpc2 "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func VideoProcessDaemon(db *mongo.Database) {
	videosRepo := repositories.NewVideosRepository(db)
	filesService := services.NewFilesService()

	useCase := usecases.NewProcessVideo(videosRepo, filesService)

	intervalTime := 5
	log.Printf("Video daemon started. Runing in: %d", intervalTime)

	for {
		useCase.Execute()
		time.Sleep(time.Duration(intervalTime) * time.Second)
	}
}

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

	if os.Getenv("RUN_DAEMON") == "true" {
		go VideoProcessDaemon(db)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
