package grpc

import (
	"context"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersServer struct {
	pb.UsersServiceServer
	createUserUseCase *usecases.CreateUserUseCase
	loginUseCase      *usecases.LoginUseCase
}

func NewUsersServer(db *mongo.Database) *UsersServer {
	usersRepository := repositories.NewUsersRepository(db)
	return &UsersServer{
		createUserUseCase: usecases.NewCreateUserUseCase(usersRepository),
		loginUseCase:      usecases.NewLoginUseCase(usersRepository),
	}
}

func (s *UsersServer) Create(_ context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user, err := s.createUserUseCase.Execute(req.Name, req.Email, req.Phone, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return entityToGrpcUser(user), nil
}

func (s *UsersServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.User, error) {
	login, err := s.loginUseCase.Execute(req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return entityToGrpcUser(login.User), nil
}

func entityToGrpcUser(user *entities.User) *pb.User {
	return &pb.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email.Email,
		Phone: user.Phone.Phone,
		Type:  user.Type.Id,
	}
}
