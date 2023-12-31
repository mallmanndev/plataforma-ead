package entities_test

import (
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	email, _ := value_objects.NewEmailAddress("matheus@email.com")
	phone, _ := value_objects.NewPhone("5599999999")
	t.Run("Should return error when user name is invalid", func(t *testing.T) {
		_, err := entities.NewUser("", email, phone, "123456789")
		assert.EqualError(t, err, "User name invalid!")
	})

	t.Run("Should return error when user password is invalid", func(t *testing.T) {
		_, err := entities.NewUser("Matheus", email, phone, "")
		assert.EqualError(t, err, "User password invalid!")

		_, err = entities.NewUser("Matheus", email, phone, "12345")
		assert.EqualError(t, err, "User password invalid!")
	})

	t.Run("Shoud return user entity when user is valid", func(t *testing.T) {
		user, err := entities.NewUser("Matheus", email, phone, "123456789")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, user.Name, "Matheus")
		assert.Equal(t, user.Email.Email, "matheus@email.com")
		assert.Equal(t, user.Phone.Phone, "5599999999")
	})
}
