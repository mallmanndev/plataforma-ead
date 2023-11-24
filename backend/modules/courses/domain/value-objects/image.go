package value_objects

type Image struct {
	url string
}

func NewImage(Url string) *Image {
	return &Image{url: Url}
}

func (i *Image) Url() string {
	return i.url
}
