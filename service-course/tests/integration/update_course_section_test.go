package integration_test

import (
	"context"
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	fixtures "github.com/matheusvmallmann/plataforma-ead/service-course/tests/fixtures/courses"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/status"
)

func TestUpdateCourseSection(t *testing.T) {
	db, closeDB := testutils.DatabaseConnection()
	ctx, client, closer := testutils.CoursesServer(db)
	courseRepo := repositories.NewCourseRepositories(db)

	defer func() {
		closeDB()
		closer()
	}()

	var setup = func(t *testing.T) {
		_, err := db.Collection("courses").InsertOne(context.TODO(), fixtures.CursoCompleto)
		if err != nil {
			t.Fatal(err)
		}
	}

	var teardown = func(t *testing.T) {
		db.Collection("courses").DeleteOne(context.Background(), bson.M{"_id": "3d515009-56eb-4ed0-aea5-182bd783085e"})
	}

	t.Run("when_course_not_found", func(t *testing.T) {
		request := &pb.UpdateCourseSectionRequest{
			UserId:      "fbf761f5-a9d8-4c39-87d6-4718cab4573b",
			Id:          "0f3e3d97-eeff-4151-a6dc-7a10d27ae6be",
			Name:        "Go Lang Course",
			Description: "This is a Go lang course",
		}

		_, err := client.UpdateSection(ctx, request)
		s, _ := status.FromError(err)
		assert.Equal(t, "Internal", s.Code().String())
		assert.Equal(t, "[Update Section] Course not found.", s.Message())
	})

	t.Run("when_user_is_different_from_instructor", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.UpdateCourseSectionRequest{
			Id:          "3d515009-56eb-4ed0-aea5-182bd783085e",
			UserId:      "fbf761f5-a9d8-4c39-87d6-4718cab4573b",
			Name:        "Go Lang Course",
			Description: "This is a Go lang course",
		}

		_, err := client.UpdateSection(ctx, request)

		s, _ := status.FromError(err)
		assert.Equal(t, "PermissionDenied", s.Code().String())
		assert.Equal(t, "Permission denied to update section.", s.Message())
	})

	t.Run("when_request_is_invalid", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.UpdateCourseSectionRequest{
			Id:          "3d515009-56eb-4ed0-aea5-182bd783085e",
			UserId:      "9111bffd-73d9-49d8-b32c-48353674dc06",
			Name:        "Go",
			Description: "This is a Go lang course",
		}

		_, err := client.UpdateSection(ctx, request)

		s, _ := status.FromError(err)
		assert.Equal(t, "InvalidArgument", s.Code().String())
		assert.Equal(t, "[Course Section] Invalid 'name': must be longer than 5.", s.Message())
	})

	t.Run("when_update_section_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.UpdateCourseSectionRequest{
			Id:          "3d515009-56eb-4ed0-aea5-182bd783085e",
			UserId:      "9111bffd-73d9-49d8-b32c-48353674dc06",
			Name:        "Go Lang Course altered",
			Description: "This is a Go lang course altered",
		}

		res, err := client.UpdateSection(ctx, request)

		assert.Nil(t, err)
		if assert.NotNil(t, res) {
			assert.Equal(t, "Go Lang Course altered", res.Sections[0].Name)
			assert.Equal(t, "This is a Go lang course altered", res.Sections[0].Description)
		}

		course, _ := courseRepo.FindById("3d515009-56eb-4ed0-aea5-182bd783085e")
		if assert.NotNil(t, course) {
			assert.Equal(t, "Go Lang Course altered", course.Sections()[0].Name())
			assert.Equal(t, "This is a Go lang course altered", course.Sections()[0].Description())
		}
	})
}
