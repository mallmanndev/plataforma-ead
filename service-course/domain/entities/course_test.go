package entities_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"testing"
)

func TestCourseEntity(t *testing.T) {
	testCenarios := []struct {
		name, description string
		image             *value_objects.Image
		instructorId      string
		expected          *entities.Course
		expectedErr       string
	}{
		{"Go", "Go", nil, "fdsfsfsfsf", nil, "Invalid course name (min: 5)!"},
		{"Go Course", "Go", nil, "fdsfsfsfsf", nil, "Invalid course description (min: 20)!"},
		{"Go Course", "This is a Go Lang Course!", nil, "fdsfsfsfsf", nil, "Invalid instructor ID!"},
	}

	for _, test := range testCenarios {
		t.Run(fmt.Sprintf("Should return exception when course is invalid (%s)", test.expectedErr), func(t *testing.T) {
			_, err := entities.NewCourse(test.name, test.description, test.image, test.instructorId)
			if err == nil {
				t.Error("Error must be not nil!")
			}
			if err.Error() != test.expectedErr {
				t.Errorf("Error must be: %s, received: %s!", test.expectedErr, err.Error())
			}
		})
	}

	t.Run("Should return course pointer when struct is valid", func(t *testing.T) {
		course, err := entities.NewCourse("Go Course", "This is a Go Lang Course!", nil, uuid.NewString())

		if err != nil {
			t.Errorf("Error must be nil, received(%s)!", err.Error())
		}
		if course == nil {
			t.Error("Course must be not nil!")
		}
		if course.IsVisible() {
			t.Error("Course must be created not visible!")
		}
	})
}
