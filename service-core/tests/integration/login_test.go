package integration_test

import (
	"context"
	"fmt"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-core/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-core/tests/integration/fixtures"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-core/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-core/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"testing"
)

func setupLoginTest(t *testing.T) (context.Context, pb.UsersServiceClient, func()) {
	db := utils.GetDb("test")
	usersRepository := repositories.NewUsersRepository(db)
	db.Exec("DELETE FROM users WHERE deleted_at IS NULL;")

	err := usersRepository.Create(fixtures.StudentUser)
	assert.Nil(t, err)

	ctx, client, closer := testutils.UserServer()

	return ctx, client, closer
}

func TestLoginGrpcRoute(t *testing.T) {
	t.Run("Should return error when email is incorrect", func(t *testing.T) {
		// GIVEN
		ctx, client, closer := setupLoginTest(t)
		defer closer()

		// WHEN
		request := &pb.LoginRequest{
			Email:    "fulano@email.com",
			Password: "87654321",
		}
		user, err := client.Login(ctx, request)
		s, _ := status.FromError(err)

		// THEN
		assert.Nil(t, user)
		if assert.NotNil(t, s) {
			assert.Equal(t, "User not found!", s.Message())
			assert.Equal(t, "InvalidArgument", s.Code().String())
		}
	})

	t.Run("Should return error when password is incorrect", func(t *testing.T) {
		// GIVEN
		ctx, client, closer := setupLoginTest(t)
		defer closer()

		// WHEN
		request := &pb.LoginRequest{
			Email:    "matheus@email.com",
			Password: "87654321",
		}
		user, err := client.Login(ctx, request)
		s, _ := status.FromError(err)

		// THEN
		assert.Nil(t, user)
		if assert.NotNil(t, s) {
			assert.Equal(t, "Invalid password!", s.Message())
			assert.Equal(t, "InvalidArgument", s.Code().String())
		}
	})

	t.Run("Should return user when email and password is not incorrect", func(t *testing.T) {
		// GIVEN
		ctx, client, closer := setupLoginTest(t)
		defer closer()

		// WHEN
		request := &pb.LoginRequest{
			Email:    "matheus@email.com",
			Password: "12345678",
		}
		user, err := client.Login(ctx, request)

		// THEN
		assert.Nil(t, err)
		if assert.NotNil(t, user) {
			fmt.Println(user)
			assert.NotNil(t, user.Id)
			assert.Equal(t, "Matheus", user.Name)
			assert.Equal(t, "559999048223", user.Phone)
			assert.Equal(t, "matheus@email.com", user.Email)
		}
	})
}
