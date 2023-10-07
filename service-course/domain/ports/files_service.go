package ports

type FileInfo struct {
	Url  string
	Size int64
	Type string
}

type FilesService interface {
	CreateFile(File FileInfo) (FilesService, error)
	SendChunk(chunk []byte) error
	Close() error
	Remove() error
}
