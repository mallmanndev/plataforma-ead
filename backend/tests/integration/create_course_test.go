package integration_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/backend/tests/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
)

func TestCreateCourse(t *testing.T) {
	db, closeDB := testutils.DatabaseConnection()
	ctx, client, closeGrpc := testutils.CoursesServer(db)

	defer func() {
		closeDB()
		closeGrpc()
	}()

	t.Run("Should return error when instructor is not provided", func(t *testing.T) {
		request := &pb.CreateCourseRequest{
			Name:        "Go Lang Course",
			Description: "This is a Go lang course",
			Instructor:  nil,
		}

		_, err := client.Create(ctx, request)
		s, _ := status.FromError(err)
		assert.Equal(t, "InvalidArgument", s.Code().String())
		assert.Equal(t, "Instructor is required.", s.Message())
	})

	t.Run("Should return error when people is invalid", func(t *testing.T) {
		request := &pb.CreateCourseRequest{
			Name:        "Go",
			Description: "",
			Instructor: &pb.People{
				Id:       "123",
				Name:     "Matheus",
				Type:     "admin",
				PhotoUrl: "",
			},
		}

		_, err := client.Create(ctx, request)
		s, _ := status.FromError(err)
		expectedStatus := "InvalidArgument"
		expectedMessage := "[People] Invalid 'id': must be valid UUID."
		if s.Code().String() != expectedStatus {
			t.Errorf("Invalid gRPC status code. (Expected: %s, Received: %s)", expectedStatus, s.Code())
		}
		if s.Message() != expectedMessage {
			t.Errorf("Invalid gRPC message. (Expected: %s, Received: %s)", expectedMessage, s.Message())
		}
	})

	t.Run("Should return error when course is invalid", func(t *testing.T) {
		request := &pb.CreateCourseRequest{
			Name:        "Go",
			Description: "",
			Instructor: &pb.People{
				Id:       uuid.NewString(),
				Name:     "Matheus",
				Type:     "admin",
				PhotoUrl: "",
			},
		}

		_, err := client.Create(ctx, request)
		s, _ := status.FromError(err)
		expectedStatus := "InvalidArgument"
		expectedMessage := "[Course] Invalid 'name': must be longer than 5."
		if s.Code().String() != expectedStatus {
			t.Errorf("Invalid gRPC status code. (Expected: %s, Received: %s)", expectedStatus, s.Code())
		}
		if s.Message() != expectedMessage {
			t.Errorf("Invalid gRPC message. (Expected: %s, Received: %s)", expectedMessage, s.Message())
		}
	})

	t.Run("Should return success when course is created successfully", func(t *testing.T) {
		request := &pb.CreateCourseRequest{
			Name:        "Go Lang Course",
			Description: "This is a Go lang course",
			Instructor: &pb.People{
				Id:       uuid.NewString(),
				Name:     "Matheus",
				Type:     "admin",
				PhotoUrl: "",
			},
		}

		response, err := client.Create(ctx, request)
		if err != nil {
			t.Error("Error must be nil!")
		}
		if response == nil {
			t.Error("Response must be not nil!")
		}
		if response.Id == "" ||
			response.Name != "Go Lang Course" ||
			response.Description != "This is a Go lang course" ||
			response.Visible != false ||
			response.CreatedAt == "" {
			t.Error("Course response are invalid!")
		}
	})
}
