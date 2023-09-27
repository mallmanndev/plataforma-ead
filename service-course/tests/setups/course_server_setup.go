package setups

import (
	"context"
	"errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	testutils "github.com/matheusvmallmann/plataforma-ead/service-course/tests/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CourseServerSetup(t *testing.T) (
	context.Context,
	pb.CoursesServiceClient,
	func(),
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
