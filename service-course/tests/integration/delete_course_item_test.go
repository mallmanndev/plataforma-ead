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

func TestDeleteCourseItem(t *testing.T) {
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

	t.Run("when_item_is_not_found", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.DeleteItemRequest{
			Id:     "fdsfsdf",
			UserId: "user_id_1",
		}

		_, err := client.DeleteItem(ctx, req)

		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "NotFound", s.Code().String())
			assert.Equal(t, "Item not found.", s.Message())
		}
	})

	t.Run("when_permission_is_denied", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.DeleteItemRequest{
			Id:     "item_id_1",
			UserId: "dsfsdfd",
		}

		_, err := client.DeleteItem(ctx, req)

		s, _ := status.FromError(err)
		if assert.Error(t, err) {
			assert.Equal(t, "NotFound", s.Code().String())
			assert.Equal(t, "Item not found.", s.Message())
		}
	})

	t.Run("should_update_item_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		req := &pb.DeleteItemRequest{
			Id:     "item_id_1",
			UserId: "user_id_1",
		}

		_, err := client.DeleteItem(ctx, req)

		assert.Nil(t, err)

		course, _ := courseRepo.FindBySectionId("section_id_1")
		assert.NotNil(t, course)
		section := course.FindSection("section_id_1")
		itens := section.Itens()

		assert.Len(t, itens, 1)
		assert.Equal(t, "item_id_2", itens[0].Id())
		assert.Equal(t, int16(1), itens[0].Order())
	})
}