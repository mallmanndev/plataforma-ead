package integration_test

import (
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCreateCourse(t *testing.T) {
	ctx, client, closer := testutils.CoursesServer()
	defer closer()

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
		if response.Course.Id == "" ||
			response.Course.Name != "Go Lang Course" ||
			response.Course.Description != "This is a Go lang course" ||
			response.Course.Visible != false ||
			response.Course.CreatedAt == "" {
			t.Error("Course response are invalid!")
		}
	})
}
