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

func TestUpdateSectionUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewUpdateSectionUseCase(mockCourseRepository)
	courseId := uuid.NewString()
	userId := uuid.NewString()
	sectionId := uuid.NewString()
	course, _ := entities.NewCourse(
		"A Go Lang course",
		"This is a Golang course",
		nil,
		userId,
	)

	t.Run("Should return error when course is not found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(nil, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   sectionId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Update Section] Course not found."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when instructor id is different of user id", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   sectionId,
			UserId:      uuid.NewString(), // Different ID
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "Permission denied to update section."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when section is not found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   sectionId,
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "Section not found."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when section data is invalid.", func(t *testing.T) {
		section, _ := entities.NewCourseSection("Section one", "This is a section one", course.Id())
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   section.Id(),
			UserId:      userId,
			Name:        "Fir",
			Description: "A test section",
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

	t.Run("Should return error when update section return error", func(t *testing.T) {
		section, _ := entities.NewCourseSection("Section one", "This is a section one", course.Id())
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().UpdateSection(section).Return(errors.New("Test"))
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   section.Id(),
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		_, err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Update Section] Could not update section: Test"
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return section when update section successfully", func(t *testing.T) {
		section, _ := entities.NewCourseSection("Section one", "This is a section one", course.Id())
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().UpdateSection(section).Return(nil)
		data := usecases.UpdateSectionDTO{
			CourseId:    courseId,
			SectionId:   section.Id(),
			UserId:      userId,
			Name:        "First Section",
			Description: "A test section",
		}
		updated, err := useCase.Execute(data)
		if err != nil {
			t.Errorf("Error must be nil!")
		}
		if updated.Name() != data.Name || updated.Description() != data.Description {
			t.Errorf("Received section different of expected.")
		}
	})
}
