package testutils

import (
	"context"
	grpc2 "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

func CoursesServer() (context.Context, pb.CoursesServiceClient, func()) {
	db, disconnect := utils.GetDb("test")

	coursesServer := grpc2.NewCourseServer(db)

	// Crie um ouvinte de buffer para comunicação de loopback
	listener := bufconn.Listen(1024 * 1024)

	// Crie um servidor gRPC
	server := grpc.NewServer()

	// Registre o serviço do servidor (CalculatorService) no servidor gRPC
	pb.RegisterCoursesServiceServer(server, coursesServer)

	// Inicie o servidor gRPC em uma rotina goroutine
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
		}
	}()

	ctx := context.Background()

	conn, err := grpc.DialContext(
		ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Erro ao criar conexão de teste: %v", err)
	}

	client := pb.NewCoursesServiceClient(conn)

	closer := func() {
		conn.Close()
		server.Stop()
		disconnect()
	}

	return ctx, client, closer
}
