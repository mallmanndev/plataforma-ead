package entities_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"testing"
)

func TestNewPeople(t *testing.T) {
	type args struct {
		Id    string
		Name  string
		Type  string
		Photo *value_objects.Image
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"Should return error whe ID is invalid",
			args{"123", "Matheus", "admin", nil},
			errors.New("Invalid ID format!"),
		},
		{
			"Should return error when name is invalid",
			args{uuid.NewString(), "Mat", "admin", nil},
			errors.New("Invalid name (min: 5)!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entities.NewPeople(tt.args.Id, tt.args.Name, tt.args.Type, tt.args.Photo)
			if err.Error() != tt.wantErr.Error() {
				t.Errorf("NewPeople() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	t.Run("Should return People pointer when People is invalid", func(t *testing.T) {
		people, _ := entities.NewPeople(uuid.NewString(), "Matheus Mallmann", "admin", nil)
		if people == nil {
			t.Error("People must be not nil!")
		}
	})
}
