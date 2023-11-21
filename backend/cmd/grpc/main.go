package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	grpc_courses "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	grpc_users "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/backend/pb"
	"github.com/matheusvmallmann/plataforma-ead/backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func VideoProcessDaemon(db *mongo.Database) {
	videosRepo := repositories.NewVideosRepository(db)
	filesService := services.NewFilesService()

	useCase := usecases.NewProcessVideoTwo(videosRepo, filesService)

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
	coursesServer := grpc_courses.NewCourseServer(db)
	fileUploadService := grpc_courses.NewFilesServer(db)
	usersService := grpc_users.NewUsersServer(db)
	pb.RegisterCoursesServiceServer(grpcServer, coursesServer)
	pb.RegisterFileUploadServiceServer(grpcServer, fileUploadService)
	pb.RegisterUsersServiceServer(grpcServer, usersService)

	if os.Getenv("RUN_DAEMON") == "true" {
		go VideoProcessDaemon(db)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error to serve gRPC!")
	}
}
