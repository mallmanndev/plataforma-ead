package testutils

import (
	"context"
	servers "github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/grpc"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

func FilesServer(db *mongo.Database) (context.Context, pb.FileUploadServiceClient, func()) {
	filesServer := servers.NewFilesServer(db)

	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterFileUploadServiceServer(server, filesServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
		}
	}()

	ctx := context.Background()

	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Erro ao criar conexão de teste: %v", err)
	}

	client := pb.NewFileUploadServiceClient(conn)

	closer := func() {
		server.Stop()
	}

	return ctx, client, closer
}
