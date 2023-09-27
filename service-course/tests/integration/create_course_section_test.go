package integration_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"testing"
)

func setupCreateSection(t *testing.T) (
	context.Context,
	pb.CoursesServiceClient, func(),
	*entities.Course,
	*repositories.CoursesRepositories,
) {
	db, disconnect := utils.GetDb("test")
	coursesRepo := repositories.NewCourseRepositories(db)
	ctx, client, closer := testutils.CoursesServer(db)
	closeAll := func() {
		disconnect()
		closer()
	}

	instructorId := "cc01cb11-7f45-4563-a6ea-bd159b6e705a"
	course, _ := entities.NewCourse(
		"Go lang course",
		"This is a go lang course",
		nil,
		instructorId,
	)

	err1 := db.Collection("people").Drop(context.Background())
	err2 := db.Collection("courses").Drop(context.Background())
	err3 := coursesRepo.Create(course)
	if err := errors.Join(err1, err2, err3); err != nil {
		assert.FailNow(t, "Error to setup test")
	}
	return ctx, client, closeAll, course, coursesRepo
}

func TestCreateCourseSectionWhenReturnErrors(t *testing.T) {
	ctx, client, closeAll, course, _ := setupCreateSection(t)
	defer closeAll()

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
				CourseId:    course.Id(),
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
			_, err := client.CreateCourseSection(ctx, test.request)
			s, _ := status.FromError(err)
			assert.Equal(t, test.expect.status, s.Code().String())
			assert.Equal(t, test.expect.message, s.Message())
		})
	}
}

func TestCreateCourseSuccessfully(t *testing.T) {
	ctx, client, closeAll, course, repo := setupCreateSection(t)
	defer closeAll()

	request := &pb.CreateCourseSectionRequest{
		CourseId:    course.Id(),
		UserId:      course.InstructorID(),
		Name:        "Go lang course",
		Description: "This is a go lang course",
	}

	created, err := client.CreateCourseSection(ctx, request)
	assert.Nil(t, err)
	if assert.NotNil(t, created) {
		assert.Len(t, created.Sections, 1)
		assert.Equal(t, created.Sections[0].Name, "Go lang course")
		assert.Equal(t, created.Sections[0].Description, "This is a go lang course")
	}

	dbCourse, _ := repo.FindById(course.Id())
	if assert.NotNil(t, dbCourse) {
		section := dbCourse.Sections()[0]
		assert.Len(t, dbCourse.Sections(), 1)
		assert.Equal(t, section.Name(), "Go lang course")
		assert.Equal(t, section.Description(), "This is a go lang course")
	}
}
