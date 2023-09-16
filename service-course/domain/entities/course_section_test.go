package entities_test

import (
	"errors"
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
		wantErr error
	}{
		{
			name:    "Should return error when name is invalid",
			args:    args{Name: "f", Description: "F", CourseId: "123"},
			wantErr: errors.New("Invalid section name (min: 5)!"),
		},
		{
			name:    "Should return error when course id is invalid",
			args:    args{Name: "First Section", Description: "", CourseId: "123"},
			wantErr: errors.New("Invalid course ID!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entities.NewCourseSection(tt.args.Name, tt.args.Description, tt.args.CourseId)
			if err.Error() != tt.wantErr.Error() {
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
