package models

import "time"

type VideoResolution struct {
	Resolution string `bson:"resolution"`
	URL        string `bson:"url"`
}

type VideoModel struct {
	Id          string            `bson:"_id"`
	Type        string            `bson:"type"`
	TmpUrl      string            `bson:"url"`
	Status      string            `bson:"status"`
	Duration    float32           `bson:"duration"`
	Size        int64             `bson:"size"`
	Resolutions []VideoResolution `bson:"resolutions"`
	CreatedAt   time.Time         `bson:"createdAt"`
	UpdatedAt   time.Time         `bson:"updatedAt"`
}
