package integration_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestUpdateCourse(t *testing.T) {
	db, disconnect := utils.GetDb("test")
	coursesRepo := repositories.NewCourseRepositories(db)
	ctx, client, closer := testutils.CoursesServer(db)
	defer (func() {
		disconnect()
		closer()
	})()

	instructorId := "cc01cb11-7f45-4563-a6ea-bd159b6e705a"
	course, _ := entities.NewCourse(
		"Go lang course",
		"This is a go lang course",
		nil,
		instructorId)

	if err := db.Collection("people").Drop(context.Background()); err != nil {
		t.Error("Error to clear people collection")
	}
	if err := db.Collection("courses").Drop(context.Background()); err != nil {
		t.Error("Error to clear courses collection")
	}
	if err := coursesRepo.Create(course); err != nil {
		t.Error("Error to create course")
	}

	t.Run("When returns error", func(t *testing.T) {

		type ExpectedErrors struct {
			status  string
			message string
		}
		cenarios := []struct {
			test           string
			id             string
			name           string
			description    string
			instructorId   string
			instructorName string
			instructorType string
			expect         ExpectedErrors
		}{{
			test:           "When course is not found",
			id:             uuid.NewString(),
			name:           "Go lang course",
			description:    "This is a go lang course",
			instructorId:   instructorId,
			instructorName: "Matheus Mallmann",
			instructorType: "admin",
			expect: ExpectedErrors{
				status:  "NotFound",
				message: "Course not found.",
			},
		}, {
			test:           "When permission is denied",
			id:             course.Id(),
			name:           "Go lang course",
			description:    "This is a go lang course",
			instructorId:   uuid.NewString(),
			instructorName: "Matheus Mallmann",
			instructorType: "admin",
			expect: ExpectedErrors{
				status:  "PermissionDenied",
				message: "Permission denied to update course.",
			},
		}}

		for _, test := range cenarios {
			t.Run(test.test, func(t *testing.T) {
				request := &pb.UpdateCourseRequest{
					CourseId:    test.id,
					Name:        test.name,
					Description: test.description,
					Instructor: &pb.People{
						Id:       test.instructorId,
						Name:     test.name,
						Type:     test.instructorType,
						PhotoUrl: "",
					},
				}
				_, err := client.Update(ctx, request)
				s, _ := status.FromError(err)
				assert.Equal(t, test.expect.status, s.Code().String())
				assert.Equal(t, test.expect.message, s.Message())
			})
		}
	})

	t.Run("Should return course when update course successfully", func(t *testing.T) {
		// GIVEN
		expectedResponse := struct {
			Id          string
			Name        string
			Description string
			Visible     bool
			CreatedAt   time.Time
		}{
			Id:          course.Id(),
			Name:        "Altered Go lang course",
			Description: "This is a altered go lang course",
			Visible:     false,
			CreatedAt:   course.CreatedAt(),
		}
		request := &pb.UpdateCourseRequest{
			CourseId:    course.Id(),
			Name:        "Altered Go lang course",
			Description: "This is a altered go lang course",
			Instructor: &pb.People{
				Id:       instructorId,
				Name:     "Altered Matheus Mallmann",
				Type:     "admin",
				PhotoUrl: "",
			},
		}

		// WHEN
		response, err := client.Update(ctx, request)
		databaseCourse, dbErr := coursesRepo.FindById(course.Id())

		// THEN
		if assert.NotNil(t, response) && assert.Nil(t, err) {
			assert.Equal(t, response.Id, expectedResponse.Id)
			assert.Equal(t, response.Name, expectedResponse.Name)
			assert.Equal(t, response.Description, expectedResponse.Description)
			assert.Equal(t, response.Visible, expectedResponse.Visible)
			assert.NotNil(t, response.CreatedAt)
			assert.NotNil(t, response.UpdatedAt)
		}
		if assert.NotNil(t, databaseCourse) && assert.Nil(t, dbErr) {
			assert.Equal(t, databaseCourse.Id(), expectedResponse.Id)
			assert.Equal(t, databaseCourse.Name(), expectedResponse.Name)
			assert.Equal(t, databaseCourse.Description(), expectedResponse.Description)
			assert.Equal(t, databaseCourse.IsVisible(), expectedResponse.Visible)
			assert.NotNil(t, databaseCourse.UpdatedAt())
			assert.NotNil(t, databaseCourse.CreatedAt())
		}
	})
}
