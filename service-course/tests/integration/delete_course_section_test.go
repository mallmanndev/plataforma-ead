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

func TestDeleteSection(t *testing.T) {
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
		request := &pb.DeleteCourseSectionRequest{
			UserId: "fbf761f5-a9d8-4c39-87d6-4718cab4573b",
			Id:     "0f3e3d97-eeff-4151-a6dc-7a10d27ae6be",
		}

		_, err := client.DeleteSection(ctx, request)
		s, _ := status.FromError(err)
		assert.Equal(t, "Internal", s.Code().String())
		assert.Equal(t, "[Delete Section] Course not found.", s.Message())
	})

	t.Run("when_user_is_different_from_instructor", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.DeleteCourseSectionRequest{
			UserId: "fbf761f5-a9d8-4c39-87d6-4718cab4573b",
			Id:     "3d515009-56eb-4ed0-aea5-182bd783ewfwe085e",
		}

		_, err := client.DeleteSection(ctx, request)

		s, _ := status.FromError(err)
		assert.Equal(t, "PermissionDenied", s.Code().String())
		assert.Equal(t, "Permission denied to update section.", s.Message())
	})

	t.Run("when_update_section_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.DeleteCourseSectionRequest{
			UserId: "9111bffd-73d9-49d8-b32c-48353674dc06",
			Id:     "3d515009-56eb-4ed0-aea5-182bd783ewfwe085e",
		}

		_, err := client.DeleteSection(ctx, request)

		assert.Nil(t, err)

		course, _ := courseRepo.FindById("3d515009-56eb-4ed0-aea5-182bd783085e")
		assert.NotNil(t, course)
		assert.Equal(t, 1, len(course.Sections()))
	})
}
