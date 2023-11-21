package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/tests/mocks"
)

func TestCreateCourseUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)

	useCase := usecases.NewCreateCourseUseCase(mockCourseRepository)

	t.Run("Should return error when course name is invalid", func(t *testing.T) {

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go",
				Description: "",
				UserId:      uuid.NewString(),
			})

		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "[Course] Invalid 'name': must be longer than 5."
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when people repository returns error", func(t *testing.T) {
		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				UserId:      uuid.NewString(),
			})

		expectedError := "[Create User] Could not insert or update people: test"
		if err == nil {
			t.Error("Error must not be nil")
		}
		if err.Error() != expectedError {
			t.Errorf("Expected error: %s, Received: %s.", expectedError, err.Error())
		}
	})

	t.Run("Should return error when course repository returns error", func(t *testing.T) {
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(errors.New("test"))

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				UserId:      uuid.NewString(),
			})

		expectedError := "[Create User] Could not create course: test"
		if err == nil {
			t.Error("Error must not be nil")
		}
		if err.Error() != expectedError {
			t.Errorf("Expected error: %s, Received: %s.", expectedError, err.Error())
		}
	})

	t.Run("Should return course when create successfully", func(t *testing.T) {
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(nil)

		create, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				UserId:      uuid.NewString(),
			})

		if err != nil {
			t.Error("Error must be nil!")
		}
		if create == nil {
			t.Error("Course must be not nil!")
		}
		if len(create.Id()) < 1 || create.CreatedAt().IsZero() {
			t.Error("Invalid return course")
		}
	})
}
