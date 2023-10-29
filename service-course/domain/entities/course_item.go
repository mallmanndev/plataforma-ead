package entities

import (
	"time"

	"github.com/google/uuid"
)

type CourseItem struct {
	id          string
	title       string
	description string
	sectionId   string
	itemType    string
	videoId     string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCourseItem(
	Title string,
	Description string,
	SectionId string,
	Type string,
	VideoId string,
) *CourseItem {
	return &CourseItem{
		id:          uuid.NewString(),
		title:       Title,
		description: Description,
		sectionId:   SectionId,
		itemType:    Type,
		videoId:     VideoId,
		createdAt:   time.Now(),
	}
}

func NewCourseItemComplete(
	Id string,
	Title string,
	Description string,
	SectionId string,
	Type string,
	VideoId string,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) *CourseItem {
	return &CourseItem{
		id:          Id,
		title:       Title,
		description: Description,
		sectionId:   SectionId,
		itemType:    Type,
		videoId:     VideoId,
		createdAt:   CreatedAt,
		updatedAt:   UpdatedAt,
	}
}

func (c CourseItem) Id() string {
	return c.id
}

func (c CourseItem) Title() string {
	return c.title
}

func (c CourseItem) Description() string {
	return c.description
}

func (c CourseItem) SectionId() string {
	return c.sectionId
}

func (c CourseItem) Type() string {
	return c.itemType
}

func (c CourseItem) VideoId() string {
	return c.videoId
}

func (c CourseItem) CreatedAt() time.Time {
	return c.createdAt
}

func (c CourseItem) UpdatedAt() time.Time {
	return c.updatedAt
}
