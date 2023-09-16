package entities

import "time"

type CourseItem struct {
	id          string
	title       string
	description string
	sectionId   string
	itemType    string
	createdAt   time.Time
	updatedAt   time.Time
}
