package usecases_test

import (
	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-core/aplication/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
	"github.com/matheusvmallmann/plataforma-ead/service-core/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func useCaseSetup(t *testing.T) (*gomock.Controller, *usecases.LoginUseCase, *mocks.MockUsersRepository) {
	mockCtrl := gomock.NewController(t)
	mockUsersRepository := mocks.NewMockUsersRepository(mockCtrl)
	usecase := usecases.NewLoginUseCase(mockUsersRepository)
	return mockCtrl, usecase, mockUsersRepository
}

func TestLoginUseCase(t *testing.T) {
	t.Run("Should return error when email is invalid", func(t *testing.T) {
		mockCtrl, usecase, _ := useCaseSetup(t)
		defer mockCtrl.Finish()

		_, err := usecase.Execute("matheus", "123")
		if err.Error() != "Email inválido!" {
			t.Errorf("Expected: %s, Received: %s", "Email is invalid!", err.Error())
		}
	})

	t.Run("Should return error when email is not registered", func(t *testing.T) {
		mockCtrl, usecase, repo := useCaseSetup(t)
		defer mockCtrl.Finish()

		email, _ := value_objects.NewEmailAddress("matheus@email.com")

		repo.EXPECT().FindByEmail(email).Return(nil, nil)

		_, err := usecase.Execute(email.Email, "123")
		expected := "User not found!"
		if err.Error() != expected {
			t.Errorf("Expected: %s, Received: %s", expected, err.Error())
		}
	})

	t.Run("Should return error when password is invalid", func(t *testing.T) {
		mockCtrl, usecase, repo := useCaseSetup(t)
		defer mockCtrl.Finish()

		email, _ := value_objects.NewEmailAddress("matheus@email.com")
		phone, _ := value_objects.NewPhone("5599999999")
		userType := entities.NewUserType("student", "Student")
		password := "12345678"
		user, _ := entities.NewUser("Matheus", email, phone, userType, password)

		repo.EXPECT().FindByEmail(email).Return(user, nil)

		_, err := usecase.Execute(email.Email, "123")
		expected := "Invalid password!"
		if err == nil {
			t.Errorf("Error is nil!")
		}
		if err == nil || err.Error() != expected {
			t.Errorf("Expected: %s, Received: %s", expected, err.Error())
		}
	})

	t.Run("Should return user and token when email and password is valid", func(t *testing.T) {
		mockCtrl, usecase, repo := useCaseSetup(t)
		defer mockCtrl.Finish()

		email, _ := value_objects.NewEmailAddress("matheus@email.com")
		phone, _ := value_objects.NewPhone("5599999999")
		userType := entities.NewUserType("student", "Student")
		password := "12345678"
		user, _ := entities.NewUser("Matheus", email, phone, userType, password)

		repo.EXPECT().FindByEmail(email).Return(user, nil)

		login, err := usecase.Execute(email.Email, password)
		if assert.Nil(t, err) {
			assert.NotNil(t, login.Token)
			assert.Equal(t, user.Id, login.User.Id)
			assert.Equal(t, "Matheus", login.User.Name)
			assert.Equal(t, "matheus@email.com", login.User.Email.Email)
			assert.Equal(t, "5599999999", login.User.Phone.Phone)
			assert.Equal(t, "student", login.User.Type.Id)
		}
	})
}
