package usecases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"testing"
)

func TestCreateCourseUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)

	useCase := usecases.NewCreateCourseUseCase(mockCourseRepository)

	t.Run("Should return error when course is invalid", func(t *testing.T) {
		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{"Go", "", uuid.NewString()})

		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "Invalid course name (min: 5)!"
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when not create course", func(t *testing.T) {
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(errors.New("Test!"))

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				"Go Lang Course",
				"This is a Go Lang course",
				uuid.NewString(),
			})

		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "Test!"
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return course when create successfully", func(t *testing.T) {
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(nil)

		create, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				"Go Lang Course",
				"This is a Go Lang course",
				uuid.NewString(),
			})

		if err != nil {
			t.Error("Error must be nil!")
		}
		if create == nil {
			t.Error("Course must be not nil!")
		}
	})
}
