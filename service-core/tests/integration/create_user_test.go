package integration_test

import (
	"fmt"
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-core/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-core/utils"
	"testing"
)

func TestCreateUserGrpcRoute(t *testing.T) {
	ctx, client, closer := testutils.UserServer()
	defer closer()

	t.Run("Should create user successfully", func(t *testing.T) {
		setup()
		request := &pb.CreateUserRequest{
			Name:     "Matheus",
			Email:    "matheus@email.com",
			Phone:    "559999048223",
			Password: "12345678",
		}
		response, err := client.Create(ctx, request)
		fmt.Println(response, err)
		if response.Status != 200 {
			t.Errorf("Invalid status! Expected %d, Received %d", 200, response.Status)
		}
		if response.Message != "Usuário criado com sucesso!" {
			t.Errorf("Invalid status! Expected %s, Received %s",
				"Usuário criado com sucesso!",
				response.Message)
		}
	})

	t.Run("Should return email already registered error", func(t *testing.T) {
		setup()

		request1 := &pb.CreateUserRequest{
			Name:     "Matheus",
			Email:    "matheus@email.com",
			Phone:    "559999048223",
			Password: "12345678",
		}
		client.Create(ctx, request1)

		request := &pb.CreateUserRequest{
			Name:     "Matheus",
			Email:    "matheus@email.com",
			Phone:    "559999048223",
			Password: "12345678",
		}
		response, err := client.Create(ctx, request)
		fmt.Println(response, err)
		if response.Status != 400 {
			t.Errorf("Invalid status! Expected %d, Received %d", 400, response.Status)
		}
		if response.Message != "Email already registered!" {
			t.Errorf("Invalid Message! Expected %s, Received %s",
				"Email already registered!",
				response.Message)
		}
	})
}

func setup() {
	db := utils.GetDb("test")
	db.Exec("DELETE FROM users WHERE deleted_at IS NULL;")
}
