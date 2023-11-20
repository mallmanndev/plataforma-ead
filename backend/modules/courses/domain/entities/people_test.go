package entities_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
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
		wantErr string
	}{
		{
			"Should return error whe ID is invalid",
			args{"123", "Matheus", "admin", nil},
			"[People] Invalid 'id': must be valid UUID.",
		},
		{
			"Should return error when name is invalid",
			args{uuid.NewString(), "Mat", "admin", nil},
			"[People] Invalid 'name': must be longer than 5.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entities.NewPeople(tt.args.Id, tt.args.Name, tt.args.Type, tt.args.Photo)
			if err.Error() != tt.wantErr {
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
