package grpc

import (
	"context"
	"database/sql"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
)

type UsersServer struct {
	pb.UsersServiceServer
	createUserUseCase *usecases.CreateUserUseCase
}

func NewUsersServer(db *sql.DB) *UsersServer {
	usersRepository := repositories.NewUsersRepository(db)
	useCase := usecases.NewCreateUserUseCase(usersRepository)
	return &UsersServer{createUserUseCase: useCase}
}

func (s *UsersServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUSerResponse, error) {
	res := &pb.CreateUSerResponse{}

	_, err := s.createUserUseCase.Execute(req.Name, req.Email, req.Phone, req.Password)

	if err != nil {
		res.Status = 400
		res.Message = err.Error()
		return res, nil
	}

	res.Status = 200
	res.Message = "Usuário criado com sucesso!"
	return res, nil
}
