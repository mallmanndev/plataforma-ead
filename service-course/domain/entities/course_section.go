package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type CourseSection struct {
	id          string
	name        string
	description string
	courseId    string
	itens       []*CourseItem
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCourseSection(Name string, Description string, CourseId string) (*CourseSection, error) {
	section := &CourseSection{
		id:          uuid.NewString(),
		name:        Name,
		description: Description,
		courseId:    CourseId,
		createdAt:   time.Now(),
	}

	if err := section.Validate(); err != nil {
		return nil, err
	}

	return section, nil
}

func (cs *CourseSection) Validate() error {
	if len(cs.name) < 5 {
		return errors.New("Invalid section name (min: 5)!")
	}
	if _, err := uuid.Parse(cs.courseId); err != nil {
		return errors.New("Invalid course ID!")
	}

	return nil
}
