package testutils

import (
	"context"
	"log"
	"net"

	grpc2 "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func CoursesServer(db *mongo.Database) (context.Context, pb.CoursesServiceClient, func()) {
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
		ctx, "bufnet", grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Erro ao criar conexão de teste: %v", err)
	}

	client := pb.NewCoursesServiceClient(conn)

	closer := func() {
		conn.Close()
		server.Stop()
	}

	return ctx, client, closer
}
