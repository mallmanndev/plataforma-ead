package value_objects_test

import (
	"testing"

	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"github.com/stretchr/testify/require"
)

func TestNewUrl(t *testing.T) {
	t.Run("should create a valid url", func(t *testing.T) {
		url, err := value_objects.NewUrl("https://www.google.com")
		require.Nil(t, err)
		require.NotNil(t, url)
	})

	t.Run("should_return_an_error_for_invalid_url", func(t *testing.T) {
		url, err := value_objects.NewUrl("http//invalid-url.com")

		require.NotNil(t, err)
		require.Nil(t, url)
	})
}
