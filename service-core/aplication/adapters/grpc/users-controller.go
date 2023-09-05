package grpc

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
)

type UsersServer struct {
	pb.UsersServiceServer
}

func (s *UsersServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUSerResponse, error) {
	res := &pb.CreateUSerResponse{}

	usersRepository := repositories.NewUsersRepository()
	createUserUseCase := usecases.NewCreateUserUseCase(usersRepository)

	user, err := createUserUseCase.Execute(req.GetName(), req.GetEmail(), req.GetPhone(), req.GetPassword())
	if err != nil {
		res.Status = 400
		res.Message = "Não foi possível cadastrar o usuário!"
		return res, nil
	}

	res.Status = 200
	res.Message = "Usuário criado com sucesso! Id:" + user.Id
	return res, nil
}
