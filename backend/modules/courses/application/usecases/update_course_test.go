package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/tests/mocks"
)

func TestUpdateCourseUseCase(t *testing.T) {

	t.Run("Should return error when course name is invalid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
		useCase := usecases.NewUpdateCourseUseCase(mockCourseRepository)

		courseId := uuid.NewString()
		instructorId := uuid.NewString()
		course := entities.NewCourseComplete(
			courseId,
			"Go Lang Course",
			"This is a Go Lang course",
			nil,
			instructorId,
			true,
			time.Now(),
			time.Now(),
			"",
		)
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		_, err := useCase.Execute(
			usecases.UpdateCourseUseCaseDTO{
				Id:          courseId,
				Name:        "Go",
				Description: "",
				UserId:      instructorId,
			})
		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "[Course] Invalid 'name': must be longer than 5."
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when not find course", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
		useCase := usecases.NewUpdateCourseUseCase(mockCourseRepository)

		courseId := uuid.NewString()
		mockCourseRepository.EXPECT().FindById(courseId).Return(nil, nil)
		_, err := useCase.Execute(
			usecases.UpdateCourseUseCaseDTO{
				Id:          courseId,
				Name:        "Go",
				Description: "",
				UserId:      uuid.NewString(),
			})
		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "Course not found."
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when courses repository return error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
		useCase := usecases.NewUpdateCourseUseCase(mockCourseRepository)

		courseId := uuid.NewString()
		instructorId := uuid.NewString()
		course := entities.NewCourseComplete(
			courseId,
			"Go Lang Course",
			"This is a Go Lang course",
			nil,
			instructorId,
			true,
			time.Now(),
			time.Now(),
			"",
		)
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().Update(gomock.Any()).Return(errors.New("Test"))
		_, err := useCase.Execute(
			usecases.UpdateCourseUseCaseDTO{
				Id:          courseId,
				Name:        "Javascript course",
				Description: "This is a Javascript course",
				UserId:      instructorId,
			})
		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "[Update Course] Could not update course: Test"
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when course instructor is difference of user", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
		useCase := usecases.NewUpdateCourseUseCase(mockCourseRepository)

		courseId := uuid.NewString()
		instructorId := uuid.NewString()
		course := entities.NewCourseComplete(
			courseId,
			"Go Lang Course",
			"This is a Go Lang course",
			nil,
			instructorId,
			true,
			time.Now(),
			time.Now(),
			"",
		)
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		_, err := useCase.Execute(
			usecases.UpdateCourseUseCaseDTO{
				Id:          courseId,
				Name:        "Javascript course",
				Description: "This is a Javascript course",
				UserId:      uuid.NewString(),
			})
		if err == nil {
			t.Error("Error must be not nil!")
		}
		expectedErr := "Permission denied to update course."
		if err.Error() != expectedErr {
			t.Errorf("Ivalid error! Expected: %s, Received: %s.", expectedErr, err.Error())
		}
	})

	t.Run("Should update course successfully", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)
		useCase := usecases.NewUpdateCourseUseCase(mockCourseRepository)

		courseId := uuid.NewString()
		instructorId := uuid.NewString()
		course := entities.NewCourseComplete(
			courseId,
			"Go Lang Course",
			"This is a Go Lang course",
			nil,
			instructorId,
			true,
			time.Now(),
			time.Now(),
			"",
		)
		mockCourseRepository.EXPECT().FindById(courseId).Return(course, nil)
		mockCourseRepository.EXPECT().Update(gomock.Any()).Return(nil)
		editedCourse, err := useCase.Execute(
			usecases.UpdateCourseUseCaseDTO{
				Id:          courseId,
				Name:        "Javascript course",
				Description: "This is a Javascript course",
				UserId:      instructorId,
			})
		if err != nil {
			t.Errorf("Error must be nil. Error: %s", err.Error())
		}
		if editedCourse.Name() != "Javascript course" || editedCourse.Description() != "This is a Javascript course" {
			t.Error("Course not edited!")
		}
	})

}
