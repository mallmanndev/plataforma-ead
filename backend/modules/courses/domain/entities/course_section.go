package entities

import (
	"time"

	"github.com/google/uuid"
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
)

type CourseSection struct {
	id          string
	name        string
	description string
	courseId    string
	itens       []*CourseItem
	order       int16
	avaliation  *value_objects.Url
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCourseSection(
	Name string,
	Description string,
	CourseId string,
	avaliation *value_objects.Url,
) (*CourseSection, error) {
	section := &CourseSection{
		id:          uuid.NewString(),
		name:        Name,
		description: Description,
		courseId:    CourseId,
		avaliation:  avaliation,
		createdAt:   time.Now(),
	}
	if err := section.Validate(); err != nil {
		return nil, err
	}
	return section, nil
}

type NewCompleteSectionData struct {
	Id          string
	Name        string
	Description string
	CourseId    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Avaliation  *value_objects.Url
}

func NewCompleteSection(data NewCompleteSectionData) *CourseSection {
	return &CourseSection{
		id:          data.Id,
		name:        data.Name,
		description: data.Description,
		courseId:    data.CourseId,
		createdAt:   data.CreatedAt,
		updatedAt:   data.UpdatedAt,
		avaliation:  data.Avaliation,
	}
}

func (cs *CourseSection) Update(
	Name string,
	Description string,
	avaliation *value_objects.Url,
) error {
	cs.name = Name
	cs.description = Description
	cs.avaliation = avaliation
	return cs.Validate()
}

func (cs *CourseSection) Validate() error {
	if len(cs.name) < 5 {
		return errs.NewInvalidAttributeError(
			"Course Section",
			"name",
			"must be longer than 5")
	}
	return nil
}

func (cs *CourseSection) Id() string {
	return cs.id
}

func (cs *CourseSection) Name() string {
	return cs.name
}

func (cs *CourseSection) Description() string {
	return cs.description
}

func (cs *CourseSection) CourseId() string {
	return cs.courseId
}

func (cs *CourseSection) Itens() []*CourseItem {
	return cs.itens
}

func (cs *CourseSection) Order() int16 {
	return cs.order
}

func (cs *CourseSection) Avaliation() *value_objects.Url {
	return cs.avaliation
}

func (cs *CourseSection) SetOrder(order int16) *CourseSection {
	cs.order = order
	return cs
}

func (cs *CourseSection) CreatedAt() time.Time {
	return cs.createdAt
}

func (cs *CourseSection) UpdatedAt() time.Time {
	return cs.updatedAt
}

func (cs *CourseSection) AddItem(item *CourseItem) {
	newOrder := len(cs.itens) + 1
	item.SetOrder(int16(newOrder))
	cs.itens = append(cs.itens, item)
}

func (cs *CourseSection) RemoveItem(Id string) {
	newItens := make([]*CourseItem, 0, len(cs.itens)-1)
	for _, item := range cs.itens {
		if item.id != Id {
			newItens = append(newItens, item)
		}
	}
	cs.itens = newItens
}
