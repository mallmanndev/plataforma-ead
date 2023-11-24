package entities_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestNewCourseSection(t *testing.T) {
	type args struct {
		Name        string
		Description string
		CourseId    string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.CourseSection
		wantErr string
	}{
		{
			name:    "Should return error when name is invalid",
			args:    args{Name: "f", Description: "F", CourseId: "123"},
			wantErr: "[Course Section] Invalid 'name': must be longer than 5.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avaliation, _ := value_objects.NewUrl("http://www.google.com")
			_, err := entities.NewCourseSection(tt.args.Name, tt.args.Description, tt.args.CourseId, avaliation)
			assert.ErrorContains(t, err, tt.wantErr)
		})
	}

	t.Run("Should return session when is valid", func(t *testing.T) {
		// GIVEN
		avaliation, _ := value_objects.NewUrl("http://www.google.com")

		// WHEN
		section, err := entities.NewCourseSection("First Section", "", uuid.NewString(), avaliation)

		// THEN
		assert.Nil(t, err)
		assert.NotNil(t, section)
	})
}
