package entities_test

import (
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"testing"
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
			_, err := entities.NewCourseSection(tt.args.Name, tt.args.Description, tt.args.CourseId)
			if err.Error() != tt.wantErr {
				t.Errorf("NewCourseSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	t.Run("Should return session when is valid", func(t *testing.T) {
		got, err := entities.NewCourseSection("First Section", "", uuid.NewString())
		if err != nil {
			t.Error("Error must be nil!")
		}
		if got == nil {
			t.Error("Section must be not nil")
		}
	})
}
