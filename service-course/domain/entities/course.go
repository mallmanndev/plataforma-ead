package entities

import (
	"github.com/google/uuid"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
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

func NewCourseComplete(
	Id string, Name string, Description string, Image *value_objects.Image,
	InstructorId string, Visible bool, CreatedAt time.Time, UpdatedAt time.Time,
) *Course {
	return &Course{
		id: Id, name: Name, description: Description, image: Image, instructorID: InstructorId,
		visible: Visible, createdAt: CreatedAt, updatedAt: UpdatedAt,
	}
}

func (c *Course) Validate() error {
	if len(c.name) < 5 {
		return errs.NewInvalidAttributeError(
			"Course",
			"name",
			"must be longer than 5")
	}
	if len(c.description) < 20 {
		return errs.NewInvalidAttributeError(
			"Course",
			"description",
			"must be longer than 20")
	}

	return nil
}

func (c *Course) Update(Name string, Description string) error {
	c.name = Name
	c.description = Description
	return c.Validate()
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

func (c *Course) FindSection(Id string) *CourseSection {
	for _, valor := range c.sections {
		if valor.id == Id {
			return valor
		}
	}
	return nil
}
