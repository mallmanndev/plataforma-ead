package errs_test

import (
	"errors"
	"fmt"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGrpcParser(t *testing.T) {
	t.Run("When error is 'InvalidAttributeError'", func(t *testing.T) {
		err := errs.NewInvalidAttributeError(
			"Users",
			"name",
			"must be longer than 5")
		grpcError := errs.NewGrpcError(err)
		s, ok := status.FromError(grpcError)
		if !ok {
			t.Errorf("Invalid error!")
		}
		if s.Code() != codes.InvalidArgument {
			t.Errorf("Expected status code: %s, Received: %s.", codes.InvalidArgument, s.Code())
		}
		if s.Message() != err.Error() {
			t.Errorf("Expected message: %s, Received: %s.", s.Message(), err.Error())
		}
	})

	t.Run("When error is 'UseCaseError'", func(t *testing.T) {
		err := errs.NewCreateUserUseCaseError(
			"Could not insert or update people", nil)
		grpcError := errs.NewGrpcError(err)
		s, ok := status.FromError(grpcError)
		if !ok {
			t.Errorf("Invalid error!")
		}
		if s.Code() != codes.Internal {
			t.Errorf("Expected status code: '%s', Received: '%s'.", codes.InvalidArgument, s.Code())
		}
		expectedErrorMessage := "[Create User] Could not insert or update people."
		if s.Message() != expectedErrorMessage {
			t.Errorf("Expected message: '%s', Received: '%s'.", s.Message(), expectedErrorMessage)
		}
	})

	t.Run("When error is 'UseCaseError' ant have wrap error", func(t *testing.T) {
		err := errs.NewCreateUserUseCaseError(
			"Could not insert or update people", errors.New("Test!"))
		grpcError := errs.NewGrpcError(err)
		s, ok := status.FromError(grpcError)
		if !ok {
			t.Errorf("Invalid error!")
		}
		if s.Code() != codes.Internal {
			t.Errorf("Expected status code: %s, Received: %s.", codes.InvalidArgument, s.Code())
		}
		expectedErrorMessage := "[Create User] Could not insert or update people."
		fmt.Println(s.Message())
		if s.Message() != expectedErrorMessage {
			t.Errorf("Expected message: %s, Received: %s.", s.Message(), expectedErrorMessage)
		}
	})
}
