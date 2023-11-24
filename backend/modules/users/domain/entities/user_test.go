package entities_test

import (
	"testing"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	studentUserType := &entities.UserType{
		Id:   "student",
		Name: "Student",
	}
	email, _ := value_objects.NewEmailAddress("matheus@email.com")
	phone, _ := value_objects.NewPhone("5599999999")
	t.Run("Should return error when user name is invalid", func(t *testing.T) {
		_, err := entities.NewUser("", email, phone, studentUserType, "123456789")
		assert.EqualError(t, err, "User name invalid!")
	})

	t.Run("Should return error when user password is invalid", func(t *testing.T) {
		_, err := entities.NewUser("Matheus", email, phone, studentUserType, "")
		assert.EqualError(t, err, "User password invalid!")

		_, err = entities.NewUser("Matheus", email, phone, studentUserType, "12345")
		assert.EqualError(t, err, "User password invalid!")
	})

	t.Run("Shoud return user entity when user is valid", func(t *testing.T) {
		user, err := entities.NewUser("Matheus", email, phone, studentUserType, "123456789")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, user.Name, "Matheus")
	})
}
