package usecases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
	"github.com/matheusvmallmann/plataforma-ead/service-core/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserUseCase(t *testing.T) {
	t.Run("Should return error when user is invalid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)

		useCase := usecases.NewCreateUserUseCase(mockUsersRepository)
		user, err := useCase.Execute("Matheus", "matheus@email.com", "55999999999", "123456")
		assert.Nil(t, user)
		assert.EqualError(t, err, "User password invalid!")
	})

	t.Run("Should return error when user email is already registered", func(t *testing.T) {
		// GIVEN
		email := "matheus@email.com"
		emailVO, _ := value_objects.NewEmailAddress(email)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)
		mockUsersRepository.EXPECT().FindByEmail(emailVO).
			Return(&entities.User{Email: emailVO}, nil)

		useCase := usecases.NewCreateUserUseCase(mockUsersRepository)
		user, err := useCase.Execute("Matheus", email, "55999999999", "123456789")
		assert.Nil(t, user)
		assert.EqualError(t, err, "Email already registered!")
	})

	t.Run("Should return error when not find user by email", func(t *testing.T) {
		// GIVEN
		email := "matheus@email.com"

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)
		mockUsersRepository.EXPECT().FindByEmail(gomock.Any()).
			Return(nil, errors.New("Test!"))

		useCase := usecases.NewCreateUserUseCase(mockUsersRepository)
		user, err := useCase.Execute("Matheus", email, "55999999999", "123456789")
		assert.Nil(t, user)
		assert.EqualError(t, err, "Test!")
	})

	t.Run("Should return error when not create user", func(t *testing.T) {
		// GIVEN
		email := "matheus@email.com"

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)
		mockUsersRepository.EXPECT().FindByEmail(gomock.Any()).
			Return(nil, nil)
		mockUsersRepository.EXPECT().Create(gomock.Any()).Return(errors.New("Test!"))

		useCase := usecases.NewCreateUserUseCase(mockUsersRepository)
		user, err := useCase.Execute("Matheus", email, "55999999999", "123456789")
		assert.Nil(t, user)
		assert.EqualError(t, err, "Test!")
	})

	t.Run("Should register password successfully", func(t *testing.T) {
		// GIVEN
		email := "matheus@email.com"
		emailVO, _ := value_objects.NewEmailAddress(email)

		phone := "55999999999"
		phoneVO, _ := value_objects.NewPhone(phone)

		expectedUser := &entities.User{
			Name:  "Matheus",
			Email: emailVO,
			Phone: phoneVO,
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)
		mockUsersRepository.EXPECT().FindByEmail(gomock.Any()).
			Return(nil, nil)
		mockUsersRepository.EXPECT().Create(gomock.Any()).Return(nil)

		useCase := usecases.NewCreateUserUseCase(mockUsersRepository)
		user, err := useCase.Execute(expectedUser.Name, email, phone, "123456789")
		assert.NotNil(t, user.Id)
		assert.Equal(t, user.Name, expectedUser.Name)
		assert.Nil(t, err)
	})
}
