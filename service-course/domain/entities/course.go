package entities

import (
	"time"

	"github.com/google/uuid"
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
)

type Course struct {
	id          string
	name        string
	description string
	image       *value_objects.Image
	userId      string
	sections    []*CourseSection
	visible     bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCourse(Name string, Description string, Image *value_objects.Image, userId string) (*Course, error) {
	course := &Course{
		id:          uuid.NewString(),
		name:        Name,
		description: Description,
		image:       Image,
		userId:      userId,
		visible:     false,
		createdAt:   time.Now(),
	}
	if err := course.Validate(); err != nil {
		return nil, err
	}
	return course, nil
}

func NewCourseComplete(
	Id string, Name string, Description string, Image *value_objects.Image,
	userId string, Visible bool, CreatedAt time.Time, UpdatedAt time.Time,
) *Course {
	return &Course{
		id: Id, name: Name, description: Description, image: Image, userId: userId,
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
	c.updatedAt = time.Now()
	return c.Validate()
}

func (c *Course) AddSection(Section *CourseSection) {
	c.sections = append(c.sections, Section)
}

func (c *Course) RemoveSection(Id string) error {
	findKey := -1
	for key, section := range c.sections {
		if section.Id() == Id {
			findKey = key
		}
	}
	if findKey < 0 {
		return errs.NewNotFoundError("Section")
	}
	c.sections = append(c.sections[:findKey], c.sections[findKey+1:]...)
	return nil
}

func (c *Course) FindSection(Id string) *CourseSection {
	for _, valor := range c.sections {
		if valor.id == Id {
			return valor
		}
	}
	return nil
}

func (c *Course) FindItem(Id string) (*CourseItem, *CourseSection) {
	for _, section := range c.sections {
		for _, item := range section.itens {
			if item.id == Id {
				return item, section
			}
		}
	}
	return nil, nil
}

func (c *Course) ChangeOrder(SectionId string, NewOrder int) error {
	if NewOrder < 1 || NewOrder > len(c.sections) {
		return errs.NewInvalidAttributeError("Course", "order", "must be valid")
	}
	section := c.FindSection(SectionId)
	if section == nil {
		return errs.NewNotFoundError("Section")
	}

	allSections := c.Sections()
	currentOrder := int(section.Order())

	if NewOrder < currentOrder {
		for order := NewOrder; order < currentOrder; order++ {
			i := order - 1
			j := order
			iOrder := allSections[i].Order()
			jOrder := allSections[j].Order()
			allSections[i].SetOrder(jOrder)
			allSections[j].SetOrder(iOrder)
			allSections[i], allSections[j] = allSections[j], allSections[i]
		}
	}

	if NewOrder > currentOrder {
		for order := currentOrder; order < NewOrder; order++ {
			i := order - 1
			j := order
			iOrder := allSections[i].Order()
			jOrder := allSections[j].Order()
			allSections[i].SetOrder(jOrder)
			allSections[j].SetOrder(iOrder)
			allSections[i], allSections[j] = allSections[j], allSections[i]
		}
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

func (c *Course) UserId() string {
	return c.userId
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
