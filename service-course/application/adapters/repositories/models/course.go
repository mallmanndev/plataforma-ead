package models

import "time"

type CourseSectionModel struct {
	Id          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

type CourseModel struct {
	Id           string               `bson:"_id"`
	Name         string               `bson:"name"`
	Description  string               `bson:"description"`
	InstructorId string               `bson:"instructorId"`
	Visible      bool                 `bson:"visible"`
	Sections     []CourseSectionModel `bson:"sections"`
	CreatedAt    time.Time            `bson:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt"`
}
