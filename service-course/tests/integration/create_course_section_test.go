package integration_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	fixtures "github.com/matheusvmallmann/plataforma-ead/service-course/tests/fixtures/courses"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/status"
)

func TestCreateCourseSectionWhenReturnErrors(t *testing.T) {
	db, closeDB := testutils.DatabaseConnection()
	ctx, client, closer := testutils.CoursesServer(db)
	coursesRepo := repositories.NewCourseRepositories(db)

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

	type ExpectedErrors struct {
		status  string
		message string
	}
	scenarios := map[string]struct {
		request *pb.CreateCourseSectionRequest
		expect  ExpectedErrors
	}{
		"when_course_is_not_founs": {
			request: &pb.CreateCourseSectionRequest{
				CourseId:    uuid.NewString(),
				UserId:      uuid.NewString(),
				Name:        "Go lang course",
				Description: "This is a go lang course",
			},
			expect: ExpectedErrors{
				status:  "NotFound",
				message: "Course not found.",
			},
		},
		"when_permission_is_denied": {
			request: &pb.CreateCourseSectionRequest{
				CourseId:    "3d515009-56eb-4ed0-aea5-182bd783085e",
				UserId:      uuid.NewString(),
				Name:        "Go lang course",
				Description: "This is a go lang course",
			},
			expect: ExpectedErrors{
				status:  "PermissionDenied",
				message: "Permission denied to create section.",
			},
		},
	}

	for name, test := range scenarios {
		t.Run(name, func(t *testing.T) {
			setup(t)
			defer teardown(t)

			_, err := client.CreateSection(ctx, test.request)
			s, _ := status.FromError(err)
			assert.Equal(t, test.expect.status, s.Code().String())
			assert.Equal(t, test.expect.message, s.Message())
		})
	}

	t.Run("should_register_section_successfully", func(t *testing.T) {
		setup(t)
		defer teardown(t)

		request := &pb.CreateCourseSectionRequest{
			CourseId:    "3d515009-56eb-4ed0-aea5-182bd783085e",
			UserId:      "9111bffd-73d9-49d8-b32c-48353674dc06",
			Name:        "Go lang course",
			Description: "This is a go lang course",
		}

		created, err := client.CreateSection(ctx, request)
		assert.Nil(t, err)
		if assert.NotNil(t, created) {
			assert.Len(t, created.Sections, 3)
			assert.Equal(t, created.Sections[2].Name, "Go lang course")
			assert.Equal(t, created.Sections[2].Description, "This is a go lang course")
		}

		dbCourse, _ := coursesRepo.FindById("3d515009-56eb-4ed0-aea5-182bd783085e")
		if assert.NotNil(t, dbCourse) {
			section := dbCourse.Sections()[2]
			assert.Len(t, dbCourse.Sections(), 3)
			assert.Equal(t, section.Name(), "Go lang course")
			assert.Equal(t, section.Description(), "This is a go lang course")
		}
	})
}
