package value_objects

type Image struct {
	url       string
	base64    string
	imageType string
	extension string
}

func NewImage() *Image {
	return &Image{}
}

func (i *Image) Base64() string {
	return i.base64
}

func (i *Image) ImageType() string {
	return i.imageType
}

func (i *Image) Extension() string {
	return i.extension
}

func (i *Image) Url() string {
	return i.url
}
