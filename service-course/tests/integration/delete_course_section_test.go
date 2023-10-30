package integration_test

import (
	"context"
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	fixtures "github.com/matheusvmallmann/plataforma-ead/service-course/tests/fixtures"
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
		db.Collection("courses").DeleteMany(context.Background(), bson.M{})
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
			UserId: "section_id_1",
			Id:     "3d515009-56eb-4ed0-aea5-182bd783ewfwe085e",
		}

		_, err := client.DeleteSection(ctx, request)

		s, _ := status.FromError(err)
		assert.Equal(t, "Internal", s.Code().String())
		assert.Equal(t, "[Delete Section] Course not found.", s.Message())
	})

	t.Run("when_update_section_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.DeleteCourseSectionRequest{
			UserId: "user_id_1",
			Id:     "section_id_1",
		}

		_, err := client.DeleteSection(ctx, request)

		assert.Nil(t, err)

		course, _ := courseRepo.FindBySectionId("section_id_1")
		assert.Nil(t, course)
	})
}
