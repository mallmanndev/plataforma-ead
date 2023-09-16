package usecases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"testing"
)

func TestCreatePeople(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPeopleRepository := mocks.NewMockPeopleRepository(mockCtrl)

	useCase := usecases.NewCreateOrUpdatePeople(mockPeopleRepository)

	t.Run("Should return error when user data is ivalid", func(t *testing.T) {
		id := uuid.NewString()

		mockPeopleRepository.EXPECT().FindById(id).Return(nil, nil)

		_, err := useCase.Execute(usecases.CreatePeopleDTO{
			Id:   id,
			Name: "Mat",
			Type: "admin",
		})
		if err == nil {
			t.Error("Error must be not nil!")
		}
	})

	t.Run("Should return error when have error to create", func(t *testing.T) {
		id := uuid.NewString()

		mockPeopleRepository.EXPECT().FindById(id).Return(nil, nil)
		mockPeopleRepository.EXPECT().Create(gomock.Any()).Return(errors.New("Test!"))

		_, err := useCase.Execute(usecases.CreatePeopleDTO{
			Id:   id,
			Name: "Matheus mallmann",
			Type: "admin",
		})
		if err == nil {
			t.Error("Error must be not nil!")
		}
		if err.Error() != "Test!" {
			t.Errorf("Invalid error! Expected %s, Received: %s.", "Test!", err.Error())
		}
	})

	t.Run("Should return nil when user is created successfully", func(t *testing.T) {
		id := uuid.NewString()

		mockPeopleRepository.EXPECT().FindById(id).Return(nil, nil)
		mockPeopleRepository.EXPECT().Create(gomock.Any()).Return(nil)

		people, err := useCase.Execute(usecases.CreatePeopleDTO{
			Id:   id,
			Name: "Matheus mallmann",
			Type: "admin",
		})
		if err != nil {
			t.Error("Error must be nil!")
		}
		if people == nil {
			t.Error("People must be not nil!")
		}
	})
}
