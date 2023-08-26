package entities_test

import (
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	t.Run("Should return error when user name is invalid", func(t *testing.T) {
		_, err := entities.NewUser("", "matheus@email.com", "5599999999", "123456789")
		assert.EqualError(t, err, "User name invalid!")
	})

	t.Run("Should return error when user email is invalid", func(t *testing.T) {
		_, err := entities.NewUser("Matheus", "", "5599999999", "123456789")
		assert.EqualError(t, err, "User email invalid!")
	})

	t.Run("Should return error when user phone is invalid", func(t *testing.T) {
		_, err := entities.NewUser("Matheus", "matheus@email.com", "", "123456789")
		assert.EqualError(t, err, "User phone invalid!")
	})

	t.Run("Should return error when user password is invalid", func(t *testing.T) {
		_, err := entities.NewUser("Matheus", "matheus@email.com", "5599999999", "")
		assert.EqualError(t, err, "User password invalid!")

		_, err = entities.NewUser("Matheus", "matheus@email.com", "5599999999", "12345")
		assert.EqualError(t, err, "User password invalid!")
	})

	t.Run("Shoud return user entity when user is valid", func(t *testing.T) {
		user, err := entities.NewUser("Matheus", "matheus@email.com", "559999999", "123456789")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, user.Name, "Matheus")
		assert.Equal(t, user.Email, "matheus@email.com")
		assert.Equal(t, user.Phone, "559999999")
	})
}
