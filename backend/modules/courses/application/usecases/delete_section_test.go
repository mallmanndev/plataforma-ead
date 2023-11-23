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
)

func TestDeleteSectionUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")
	avaliation, _ := value_objects.NewUrl("https://www.avaliation.com")

	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
	useCase := usecases.NewDeleteSectionUseCase(mockCourseRepository)
	userId := uuid.NewString()
	sectionId := uuid.NewString()
	course, _ := entities.NewCourse(
		"A Go Lang course",
		"This is a Golang course",
		nil,
		userId,
		discordUrl,
	)

	t.Run("Should return error when course is not found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindBySectionId(sectionId).Return(nil, nil)
		data := usecases.DeleteSectionDTO{
			SectionId: sectionId,
			UserId:    userId,
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Delete Section] Course not found."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when instructor id is different of user id", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindBySectionId(sectionId).Return(course, nil)
		data := usecases.DeleteSectionDTO{
			SectionId: sectionId,
			UserId:    uuid.NewString(),
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "Permission denied to update section."
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return error when delete section return error", func(t *testing.T) {
		section, _ := entities.NewCourseSection(
			"Section one",
			"This is a section one",
			course.Id(),
			avaliation,
		)
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindBySectionId(section.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(errors.New("Test"))
		data := usecases.DeleteSectionDTO{
			SectionId: section.Id(),
			UserId:    userId,
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must not be nil!")
		}
		expectedError := "[Delete Section] Could not delete section: Test"
		if err.Error() != expectedError {
			t.Errorf("Expected: %s, Received: %s", expectedError, err.Error())
		}
	})

	t.Run("Should return section when update section successfully", func(t *testing.T) {
		section, _ := entities.NewCourseSection(
			"Section one",
			"This is a section one",
			course.Id(),
			avaliation,
		)
		course.AddSection(section)

		mockCourseRepository.EXPECT().FindBySectionId(section.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Update(course).Return(nil)
		data := usecases.DeleteSectionDTO{
			SectionId: section.Id(),
			UserId:    userId,
		}
		err := useCase.Execute(data)
		if err != nil {
			t.Errorf("Error must be nil!")
		}
	})
}
