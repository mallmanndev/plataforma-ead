package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
	"github.com/matheusvmallmann/plataforma-ead/backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateSectionUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewCreateSectionUseCase(mockCourseRepository)
	courseId := uuid.NewString()
	userId := uuid.NewString()
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")
	course, _ := entities.NewCourse(
		"A Go Lang course",
		"This is a Golang course",
		nil,
		userId,
		discordUrl,
	)

	t.Run("Should return error when section is invalid", func(t *testing.T) {
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "",
			Description: "",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "[Course Section] Invalid 'name': must be longer than 5.")
		}
	})

	t.Run("Should return error when course is not found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(nil, nil)
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "Course not found.")
		}
	})

	t.Run("Should return error when instructor id is different of user id", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      uuid.NewString(), // Different ID
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "Permission denied to create section.")
		}
	})

	t.Run("Should return error when create section returns error", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(errors.New("Test!"))
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "[Create Section] Could not create section: Test!")
		}
	})

	t.Run("Should create section successfully", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(nil)
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
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
