package entities_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"github.com/stretchr/testify/assert"
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
		{"Go", "Go", nil, "fdsfsfsfsf", nil,
			"[Course] Invalid 'name': must be longer than 5."},
		{"Go Course", "Go", nil, "fdsfsfsfsf", nil,
			"[Course] Invalid 'description': must be longer than 20."},
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
		course, err := entities.NewCourse(
			"Go Course",
			"This is a Go Lang Course!",
			nil,
			uuid.NewString())
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

func TestCourseEntitySections(t *testing.T) {
	t.Run("Should return nil when not found section", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString())
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())

		course.AddSection(section)
		findSection := course.FindSection(uuid.NewString())
		if findSection != nil {
			t.Error("Find section must return nil")
		}
	})

	t.Run("Should return nil when not found section", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString())
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())

		course.AddSection(section)
		findSection := course.FindSection(section.Id())
		if findSection != section {
			t.Error("Find section must return section.")
		}
	})
}

func TestRemoveSection(t *testing.T) {
	course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString())
	section1, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())
	section2, _ := entities.NewCourseSection("Section two", "this is a section two", course.Id())
	course.AddSection(section1)
	course.AddSection(section2)

	t.Run("When section is not found", func(t *testing.T) {
		err := course.RemoveSection(uuid.NewString())
		assert.Error(t, err, "Section not found.")
	})

	t.Run("When remove section successfully", func(t *testing.T) {
		assert.Len(t, course.Sections(), 2)
		err := course.RemoveSection(section2.Id())
		assert.Nil(t, err)
		assert.Len(t, course.Sections(), 1)
	})
}
