package entities

import (
	"errors"
	"github.com/google/uuid"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"time"
)

type People struct {
	id         string
	name       string
	peopleType string
	photo      *value_objects.Image
	createdAt  time.Time
	updatedAt  time.Time
}

func NewPeople(Id string, Name string, Type string, Photo *value_objects.Image) (*People, error) {
	people := &People{
		id:         Id,
		name:       Name,
		peopleType: Type,
		photo:      Photo,
		createdAt:  time.Now(),
	}

	if err := people.Validate(); err != nil {
		return nil, err
	}

	return people, nil
}

func (p *People) Validate() error {
	if _, err := uuid.Parse(p.id); err != nil {
		return errors.New("Invalid ID format!")
	}
	if len(p.name) < 5 {
		return errors.New("Invalid name (min: 5)!")
	}
	return nil
}

func (p *People) Update(Name string, Image *value_objects.Image) error {
	p.name = Name
	p.photo = Image

	return p.Validate()
}

func (p *People) Id() string {
	return p.id
}

func (p *People) Name() string {
	return p.name
}

func (p *People) Photo() *value_objects.Image {
	return p.photo
}

func (p *People) GetType() string {
	return p.peopleType
}