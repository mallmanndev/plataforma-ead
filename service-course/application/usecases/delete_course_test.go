package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
)

func TestDeleteCourseUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCourseRepository := mocks.NewMockCourseRepository(mockCtrl)

	discordUrl, _ := value_objects.NewUrl("https://www.discord.com")
	useCase := usecases.NewDeleteCourseUseCase(mockCourseRepository)

	t.Run("Should return error when course not found", func(t *testing.T) {
		mockCourseRepository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
		data := usecases.DeleteCourseUseCaseDataDTO{
			Id:     "123",
			UserId: "123",
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must be not nil!")
		}
		expectedErr := "Course not found."
		if err.Error() != expectedErr {
			t.Errorf("Expected error: '%s', Received '%s'.", err.Error(), expectedErr)
		}
	})

	t.Run("Should return error when userid is different to course instructor id", func(t *testing.T) {
		course, _ := entities.NewCourse(
			"A Golang course",
			"A Golang course description",
			nil,
			"321",
			discordUrl,
		)

		mockCourseRepository.EXPECT().FindById("123").Return(course, nil)
		data := usecases.DeleteCourseUseCaseDataDTO{
			Id:     "123",
			UserId: "123",
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must be not nil!")
		}
		expectedErr := "Permission denied to delete course."
		if err.Error() != expectedErr {
			t.Errorf("Expected error: '%s', Received '%s'.", expectedErr, err.Error())
		}
	})

	t.Run("Should return error when userid is different to course instructor id", func(t *testing.T) {
		course, _ := entities.NewCourse(
			"A Golang course",
			"A Golang course description",
			nil,
			"321",
			discordUrl,
		)

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Delete(course.Id()).Return(errors.New("Test!"))
		data := usecases.DeleteCourseUseCaseDataDTO{
			Id:     course.Id(),
			UserId: course.UserId(),
		}
		err := useCase.Execute(data)
		if err == nil {
			t.Errorf("Error must be not nil!")
		}
		expectedErr := "[Delete Course] Could not delete course: Test!"
		if err.Error() != expectedErr {
			t.Errorf("Expected error: '%s', Received '%s'.", expectedErr, err.Error())
		}
	})

	t.Run("Should delete course successfully", func(t *testing.T) {
		course, _ := entities.NewCourse(
			"A Golang course",
			"A Golang course description",
			nil,
			"321",
			discordUrl,
		)

		mockCourseRepository.EXPECT().FindById(course.Id()).Return(course, nil)
		mockCourseRepository.EXPECT().Delete(course.Id()).Return(nil)
		data := usecases.DeleteCourseUseCaseDataDTO{
			Id:     course.Id(),
			UserId: course.UserId(),
		}
		err := useCase.Execute(data)
		if err != nil {
			t.Errorf("Error must be nil!")
		}
	})
}
