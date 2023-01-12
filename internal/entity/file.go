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
	Name   string `json:"name" db:"name"`
	Path   string `json:"pdf_file_link" db:"pdf_file_link" db:"image_file_link" json:"image_file_link"`
	Size   int64
	File   []byte
	Type   FileType
	Status FileStatus
}
