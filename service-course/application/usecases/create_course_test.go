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

	mockPeopleRepository := mocks.NewMockPeopleRepository(mockCtrl)
	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)

	useCase := usecases.NewCreateCourseUseCase(mockPeopleRepository, mockCourseRepository)

	t.Run("Should return error when instructor name is invalid", func(t *testing.T) {
		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go",
				Description: "",
				Instructor: usecases.CreateCourseInstructorDTO{
					Id:   uuid.NewString(),
					Name: "M",
					Type: "admin",
				},
			})

		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "[People] Invalid 'name': must be longer than 5."
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when course name is invalid", func(t *testing.T) {
		mockPeopleRepository.EXPECT().Upsert(gomock.Any()).Return(nil)

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go",
				Description: "",
				Instructor: usecases.CreateCourseInstructorDTO{
					Id:   uuid.NewString(),
					Name: "Matheus Mallmann",
					Type: "admin",
				},
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
		mockPeopleRepository.EXPECT().Upsert(gomock.Any()).Return(errors.New("test"))

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				Instructor: usecases.CreateCourseInstructorDTO{
					Id:   uuid.NewString(),
					Name: "Matheus Mallmann",
					Type: "admin",
				},
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
		mockPeopleRepository.EXPECT().Upsert(gomock.Any()).Return(nil)
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(errors.New("test"))

		_, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				Instructor: usecases.CreateCourseInstructorDTO{
					Id:   uuid.NewString(),
					Name: "Matheus Mallmann",
					Type: "admin",
				},
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
		mockPeopleRepository.EXPECT().Upsert(gomock.Any()).Return(nil)
		mockCourseRepository.EXPECT().Create(gomock.Any()).Return(nil)

		create, err := useCase.Execute(
			usecases.CreateCourseUseCaseDTO{
				Name:        "Go Lang",
				Description: "This is a Golang course.",
				Instructor: usecases.CreateCourseInstructorDTO{
					Id:   uuid.NewString(),
					Name: "Matheus Mallmann",
					Type: "admin",
				},
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
