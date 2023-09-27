package usecases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup(t *testing.T) (*mocks.MockCourseRepository, *usecases.UpdateSectionUseCase, *entities.Course, func()) {
	mockCtrl := gomock.NewController(t)
	closer := func() {
		mockCtrl.Finish()
	}
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewUpdateSectionUseCase(mockCourseRepository)

	userId := uuid.NewString()
	course, _ := entities.NewCourse(
		"A Go Lang course",
		"This is a Golang course",
		nil,
		userId,
	)

	return mockCourseRepository, useCase, course, closer
}

func TestUpdateSectionUseCase(t *testing.T) {

	t.Run("Should return error when course is not found", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(nil, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   uuid.NewString(),
			UserId:      uuid.NewString(),
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "[Update Section] Course not found.")
		}
	})

	t.Run("Should return error when instructor id is different of user id", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   uuid.NewString(),
			UserId:      uuid.NewString(), // Different ID
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "Permission denied to update section.")
		}
	})

	t.Run("Should return error when section is not found", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   uuid.NewString(),
			UserId:      course.InstructorID(),
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "Section not found.")
		}
	})

	t.Run("Should return error when section data is invalid.", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		section, _ := entities.NewCourseSection("Section one", "This is a section one", course.Id())
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   section.Id(),
			UserId:      course.InstructorID(),
			Name:        "Fir",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "[Course Section] Invalid 'name': must be longer than 5.")
		}
	})

	t.Run("Should return error when update section return error", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		section, _ := entities.NewCourseSection("Section one", "This is a section one", course.Id())
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(errors.New("Test"))

		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   section.Id(),
			UserId:      course.InstructorID(),
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)

		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "[Update Section] Could not update section: Test")
		}
	})

	t.Run("Should return section when update section successfully", func(t *testing.T) {
		mockCourseRepository, useCase, course, closer := setup(t)
		defer closer()

		section, _ := entities.NewCourseSection(
			"Section one",
			"This is a section one",
			course.Id(),
		)
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(nil)

		data := usecases.UpdateSectionDTO{
			CourseId:    course.Id(),
			SectionId:   section.Id(),
			UserId:      course.InstructorID(),
			Name:        "First Section",
			Description: "A test section",
		}
		response, err := useCase.Execute(data)

		assert.Nil(t, err)
		assert.Equal(t, course, response)
		assert.Equal(t, data.Name, course.Sections()[0].Name())
		assert.Equal(t, data.Description, course.Sections()[0].Description())
	})
}
