package models

import "time"

type CourseItemModel struct {
	Id          string    `bson:"_id"`
	Title       string    `bson:"name"`
	Description string    `bson:"description"`
	Type        string    `bson:"type"`
	VideoId     string    `bson:"videoId"`
	Order       int16     `bson:"order"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

type CourseSectionModel struct {
	Id            string            `bson:"_id"`
	Name          string            `bson:"name"`
	Description   string            `bson:"description"`
	Itens         []CourseItemModel `bson:"itens"`
	AvaliationUrl string            `bson:"avaliationUrl"`
	CreatedAt     time.Time         `bson:"createdAt"`
	UpdatedAt     time.Time         `bson:"updatedAt"`
}

type CourseModel struct {
	Id          string               `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	UserId      string               `bson:"userId"`
	Visible     bool                 `bson:"visible"`
	Sections    []CourseSectionModel `bson:"sections"`
	DiscordUrl  string               `bson:"discordUrl"`
	CreatedAt   time.Time            `bson:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt"`
}
