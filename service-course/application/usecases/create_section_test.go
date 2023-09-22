package usecases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"testing"
)

func TestCreateSectionUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewCreateSectionUseCase(mockCourseRepository)
	courseId := uuid.NewString()
	userId := uuid.NewString()
	course, _ := entities.NewCourse(
		"A Go Lang course",
		"This is a Golang course",
		nil,
		userId,
	)

	t.Run("Should return error when section is invalid", func(t *testing.T) {
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "",
			Description: "",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Course Section] Invalid 'name': must be longer than 5."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
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
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Create Section] Course not found."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
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
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "Permission denied to create section."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when create section returns error", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().AddSection(gomock.Any()).Return(errors.New("Test!"))
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Create Section] Could not create section: Test!"
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should create section successfully", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().AddSection(gomock.Any()).Return(nil)
		data := usecases.CreateSectionDTO{
			CourseId:    courseId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		section, err := useCase.Execute(data)
		if err != nil {
			t.Errorf("Error must be nil!")
		}
		if section.Name() != "First Section" || section.Description() != "A test section" {
			t.Errorf("Received section different of expected.")
		}
	})
}
