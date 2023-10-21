package models

import "time"

type CourseItemModel struct {
	Id          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Type        string    `bson:"type"`
	VideoId     string    `bson:"videoId"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

type CourseSectionModel struct {
	Id          string            `bson:"_id"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Itens       []CourseItemModel `bson:"itens"`
	CreatedAt   time.Time         `bson:"createdAt"`
	UpdatedAt   time.Time         `bson:"updatedAt"`
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
