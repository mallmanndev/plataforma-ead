package value_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailAdressValueObject(t *testing.T) {
	t.Run("Should returns error when email is invalid", func(t *testing.T) {
		invalidEmails := [4]string{
			"matheus",
			"matheus@email",
			"matheus.com",
			"fdsfdfsdfdffdsfdsfsdfsdfsdfsdfsdfdf@email.com",
		}

		for _, email := range invalidEmails {
			t.Run(email, func(t *testing.T) {
				_, err := NewEmailAddress(email)
				assert.EqualError(t, err, "Email inválido!")
			})
		}
	})

	t.Run("Should returns email when email is valid", func(t *testing.T) {
		validEmail := "matheus@email.com"

		email, err := NewEmailAddress(validEmail)
		assert.Nil(t, err)
		assert.Equal(t, email.Email, validEmail)
	})
}
