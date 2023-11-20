package value_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhoneValueObject(t *testing.T) {
	t.Run("Should return error when phone is invalid", func(t *testing.T) {
		_, err := NewPhone("1234")
		assert.EqualError(t, err, "Telefone inválido!")
	})

	t.Run("Should phone when phone is valid", func(t *testing.T) {
		phone, err := NewPhone("55999999999")
		assert.Nil(t, err)
		assert.Equal(t, phone.Phone, "55999999999")
	})
}
