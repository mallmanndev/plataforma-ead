package entities

import (
	"errors"
	"github.com/google/uuid"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
	"time"
)

type Course struct {
	id           string
	name         string
	description  string
	image        *value_objects.Image
	instructorID string
	sections     []*CourseSection
	visible      bool
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCourse(Name string, Description string, Image *value_objects.Image, InstructorId string) (*Course, error) {
	course := &Course{
		id:           uuid.NewString(),
		name:         Name,
		description:  Description,
		image:        Image,
		instructorID: InstructorId,
		visible:      false,
		createdAt:    time.Now(),
	}
	if err := course.Validate(); err != nil {
		return nil, err
	}

	return course, nil
}

func (c *Course) Validate() error {
	if len(c.name) < 5 {
		return errors.New("Invalid course name (min: 5)!")
	}
	if len(c.description) < 20 {
		return errors.New("Invalid course description (min: 20)!")
	}
	if _, err := uuid.Parse(c.instructorID); err != nil {
		return errors.New("Invalid instructor ID!")
	}

	return nil
}

func (c *Course) Id() string {
	return c.id
}

func (c *Course) Name() string {
	return c.name
}

func (c *Course) Description() string {
	return c.description
}

func (c *Course) Image() *value_objects.Image {
	return c.image
}

func (c *Course) InstructorID() string {
	return c.instructorID
}

func (c *Course) Sections() []*CourseSection {
	return c.sections
}

func (c *Course) IsVisible() bool {
	return c.visible
}

func (c *Course) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Course) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Course) AddSection(Section *CourseSection) {
	c.sections = append(c.sections, Section)
}
