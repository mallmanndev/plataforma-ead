package grpcroutes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
)

type UsersServer struct {
	pb.UsersServiceServer
}

func (s *UsersServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUSerResponse, error) {
	res := &pb.CreateUSerResponse{}

	user, err := entities.NewUser(req.GetName(), req.GetEmail(), req.GetPhone(), req.GetPassword())
	if err != nil {
		res.Status = 400
		res.Message = "Usuário inválido!"
		return res, nil
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))

	res.Status = 200
	res.Message = "Usuário criado com sucesso!"

	return res, nil
}
