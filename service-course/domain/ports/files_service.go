package ports

type FileInfo struct {
	Url  string
	Size int64
	Type string
}

type GetResolutionOutput struct {
	Resolution string
	Size       int64
	Type       string
}

type FilesService interface {
	CreateFile(File FileInfo) (FilesService, error)
	SendChunk(chunk []byte) error
	WriteString(content string) error
	Close() error
	Remove() error
	GetResolution(Url string) (string, error)
	ProcessVideo(InputUrl string, OutputPath string, Resolution string) error
	Delete(Url string) error
}
