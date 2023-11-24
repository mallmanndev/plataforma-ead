package integration_test

import (
	"context"
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/pb"
	fixtures "github.com/matheusvmallmann/plataforma-ead/backend/tests/fixtures"
	testutils "github.com/matheusvmallmann/plataforma-ead/backend/tests/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/status"
)

func TestCreateCourseItem(t *testing.T) {
	db, closeDB := testutils.DatabaseConnection()
	ctx, client, closer := testutils.CoursesServer(db)
	courseRepo := repositories.NewCourseRepositories(db)

	defer func() {
		closeDB()
		closer()
	}()

	var setup = func(t *testing.T) {
		if _, err := db.Collection("courses").InsertOne(context.TODO(), fixtures.CursoCompleto); err != nil {
			t.Fatal(err)
		}

		if _, err := db.Collection("videos").InsertMany(context.TODO(), fixtures.VideosBson); err != nil {
			t.Fatal(err)
		}
	}

	var teardown = func(t *testing.T) {
		db.Collection("courses").DeleteMany(context.Background(), bson.M{})
		db.Collection("videos").DeleteMany(context.Background(), bson.M{})
	}

	t.Run("when_section_is_not_found", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.CreateItemRequest{
			SectionId:   "fdsfsdf",
			UserId:      "user_id_1",
			Title:       "Test",
			Description: "Test",
			VideoId:     "video",
		}

		_, err := client.CreateItem(ctx, req)
		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "NotFound", s.Code().String())
			assert.Equal(t, "Course not found.", s.Message())
		}
	})

	t.Run("when_permission_is_denied", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.CreateItemRequest{
			SectionId:   "section_id_1",
			UserId:      "dsfsdfsdfdsf",
			Title:       "Test",
			Description: "Test",
			VideoId:     "video",
		}

		_, err := client.CreateItem(ctx, req)
		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "PermissionDenied", s.Code().String())
			assert.Equal(t, "Permission denied to create item.", s.Message())
		}
	})

	t.Run("when_video_is_not_found", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.CreateItemRequest{
			SectionId:   "section_id_1",
			UserId:      "user_id_1",
			Title:       "Test",
			Description: "Test",
			VideoId:     "sdfsdfsdf",
		}

		_, err := client.CreateItem(ctx, req)
		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "NotFound", s.Code().String())
			assert.Equal(t, "Video not found.", s.Message())
		}
	})

	t.Run("when_video_id_is_different", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.CreateItemRequest{
			SectionId:   "section_id_1",
			UserId:      "user_id_1",
			Title:       "Test",
			Description: "Test",
			VideoId:     "video_id_2",
		}

		_, err := client.CreateItem(ctx, req)
		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "NotFound", s.Code().String())
			assert.Equal(t, "Video not found.", s.Message())
		}
	})

	t.Run("should_create_item_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.CreateItemRequest{
			SectionId:   "section_id_1",
			UserId:      "user_id_1",
			Title:       "Test",
			Description: "Test description",
			VideoId:     "video_id_1",
		}

		_, err := client.CreateItem(ctx, req)
		assert.Nil(t, err)

		course, _ := courseRepo.FindBySectionId("section_id_1")
		assert.NotNil(t, course)
		section := course.FindSection("section_id_1")
		itens := section.Itens()

		assert.Len(t, itens, 3)
		assert.Equal(t, "Test", itens[2].Title())
		assert.Equal(t, "Test description", itens[2].Description())
		assert.Equal(t, int16(3), itens[2].Order())
	})
}
