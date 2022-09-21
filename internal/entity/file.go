package entity

type (
	FileStatus int
	FileType   string
)

const (
	ClientUploadInProgress FileStatus = iota
	UploadedByClient
	ClientUploadError
	StorageUploadInProgress
	UploadedToStorage
	StorageUploadError
)

const (
	Image FileType = "image"
	Other FileType = "other"
	PDF   FileType = "pdf"
)

type File struct {
	Name   string
	Path   string
	Size   int64
	file   []byte
	Type   FileType
	Status FileStatus
}
