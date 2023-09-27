package integration_test

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/setups"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"testing"
)

func TestDeleteCourse(t *testing.T) {

	t.Run("When return errors", func(t2 *testing.T) {
		ctx, client, closeAll, course, _ := setups.CourseServerSetup(t)
		defer closeAll()

		type ExpectedError struct {
			status  string
			message string
		}
		scnarios := map[string]struct {
			id          string
			userId      string
			expectedErr ExpectedError
		}{
			"when_course_is_not_found": {
				id:     "123",
				userId: "123",
				expectedErr: ExpectedError{
					status:  "NotFound",
					message: "Course not found.",
				},
			},
			"when_user_is_different_of_instructor": {
				id:     course.Id(),
				userId: "123",
				expectedErr: ExpectedError{
					status:  "PermissionDenied",
					message: "Permission denied to delete course.",
				},
			},
		}

		for name, test := range scnarios {
			t.Run(name, func(t *testing.T) {
				request := &pb.DeleteCourseRequest{
					CourseId: test.id,
					UserId:   test.userId,
				}

				_, err := client.Delete(ctx, request)
				s, _ := status.FromError(err)
				assert.Equal(t, test.expectedErr.status, s.Code().String())
				assert.Equal(t, test.expectedErr.message, s.Message())
			})
		}
	})

	t.Run("when_delete_course_successfully", func(t *testing.T) {
		ctx, client, closeAll, course, coursesRepo := setups.CourseServerSetup(t)
		defer closeAll()

		request := &pb.DeleteCourseRequest{
			CourseId: course.Id(),
			UserId:   course.InstructorID(),
		}

		del, err := client.Delete(ctx, request)
		assert.Nil(t, err)
		assert.NotNil(t, del)
		assert.Equal(t, del.Message, "Course deleted successfully.")

		findCourse, _ := coursesRepo.FindById(course.Id())
		assert.Nil(t, findCourse)
	})
}
