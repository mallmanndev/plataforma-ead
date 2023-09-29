package integration_test

import (
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-core/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-core/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
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
		user, err := client.Create(ctx, request)
		if assert.Nil(t, err) {
			assert.NotNil(t, user.Id)
			assert.Equal(t, "Matheus", user.Name)
			assert.Equal(t, "559999048223", user.Phone)
			assert.Equal(t, "matheus@email.com", user.Email)
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

		_, err := client.Create(ctx, request)
		s, _ := status.FromError(err)
		if assert.NotNil(t, s) {
			assert.Equal(t, "Email already registered!", s.Message())
		}
	})
}

func setup() {
	db := utils.GetDb("test")
	db.Exec("DELETE FROM users WHERE deleted_at IS NULL;")
}
