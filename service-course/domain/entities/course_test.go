package entities_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"github.com/stretchr/testify/assert"
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

	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")

	for _, test := range testCenarios {
		t.Run(fmt.Sprintf("Should return exception when course is invalid (%s)", test.expectedErr), func(t *testing.T) {
			_, err := entities.NewCourse(test.name, test.description, test.image, test.instructorId, discordUrl)
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
			uuid.NewString(),
			discordUrl,
		)
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
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")
	t.Run("Should return nil when not found section", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())

		course.AddSection(section)
		findSection := course.FindSection(uuid.NewString())
		if findSection != nil {
			t.Error("Find section must return nil")
		}
	})

	t.Run("Should return nil when not found section", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())

		course.AddSection(section)
		findSection := course.FindSection(section.Id())
		if findSection != section {
			t.Error("Find section must return section.")
		}
	})
}

func TestRemoveSection(t *testing.T) {
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")

	course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
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

func TestReorderSections(t *testing.T) {
	// Arrange
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")

	var generateCourse = func() (*entities.Course, []*entities.CourseSection) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
		section1, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())
		section1.SetOrder(1)
		course.AddSection(section1)
		section2, _ := entities.NewCourseSection("Section two", "this is a section two", course.Id())
		section2.SetOrder(2)
		course.AddSection(section2)
		section3, _ := entities.NewCourseSection("Section three", "this is a section three", course.Id())
		section3.SetOrder(3)
		course.AddSection(section3)
		section4, _ := entities.NewCourseSection("Section four", "this is a section four", course.Id())
		section4.SetOrder(4)
		course.AddSection(section4)
		return course, []*entities.CourseSection{section1, section2, section3, section4}
	}

	t.Run("When order is invalid, return error", func(t *testing.T) {
		course, sections := generateCourse()
		err := course.ChangeOrder(sections[0].Id(), 0)
		assert.ErrorContains(t, err, "[Course] Invalid 'order': must be valid.")

		err2 := course.ChangeOrder(sections[1].Id(), 5)
		assert.ErrorContains(t, err2, "[Course] Invalid 'order': must be valid.")
	})

	t.Run("When section id is invalid, return error", func(t *testing.T) {
		course, _ := generateCourse()
		err := course.ChangeOrder("gggggggggggggggg", 2)
		assert.ErrorContains(t, err, "Section not found.")
	})

	t.Run("order_successfully_when_new_order_is_lower_to_order", func(t *testing.T) {
		course, testSections := generateCourse()

		err := course.ChangeOrder(testSections[2].Id(), 2)
		assert.Nil(t, err)

		sections := course.Sections()

		assert.Equal(t, testSections[0].Id(), sections[0].Id())
		assert.Equal(t, int16(1), sections[0].Order())

		assert.Equal(t, testSections[2].Id(), sections[1].Id())
		assert.Equal(t, int16(2), sections[1].Order())

		assert.Equal(t, testSections[1].Id(), sections[2].Id())
		assert.Equal(t, int16(3), sections[2].Order())

		assert.Equal(t, testSections[3].Id(), sections[3].Id())
		assert.Equal(t, int16(4), sections[3].Order())
	})

	t.Run("order_successfully_when_new_order_is_lower_to_order", func(t *testing.T) {
		course, testSections := generateCourse()

		err := course.ChangeOrder(testSections[1].Id(), 3)
		assert.Nil(t, err)

		sections := course.Sections()

		assert.Equal(t, testSections[0].Id(), sections[0].Id())
		assert.Equal(t, int16(1), sections[0].Order())

		assert.Equal(t, testSections[1].Id(), sections[2].Id())
		assert.Equal(t, int16(3), sections[2].Order())

		assert.Equal(t, testSections[2].Id(), sections[1].Id())
		assert.Equal(t, int16(2), sections[1].Order())

		assert.Equal(t, testSections[3].Id(), sections[3].Id())
		assert.Equal(t, int16(4), sections[3].Order())
	})
}

func TestCourseEntity_MakeVisible(t *testing.T) {
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")

	t.Run("when_course_not_have_sections", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)

		err := course.MakeVisible()

		assert.ErrorContains(t, err, "[Course] cannot be visible without sections")
	})

	t.Run("when_course_not_have_itens", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())
		course.AddSection(section)

		err := course.MakeVisible()

		assert.ErrorContains(t, err, "Course", "[Course] cannot be visible without itens")
	})

	t.Run("should_set_visibility_to_true", func(t *testing.T) {
		course, _ := entities.NewCourse("Go lang course", "This is a Golang course", nil, uuid.NewString(), discordUrl)
		section, _ := entities.NewCourseSection("Section one", "this is a section one", course.Id())
		item := entities.NewCourseItem("Item one", "this is a item one", section.Id(), "video", "123")
		section.AddItem(item)
		course.AddSection(section)

		err := course.MakeVisible()

		assert.Nil(t, err)
		assert.True(t, course.IsVisible())
	})
}
