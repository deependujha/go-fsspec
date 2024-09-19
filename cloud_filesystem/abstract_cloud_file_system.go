package cloud_filesystem

type AbstractCloudFileStorage interface {
	// initialize
	Init(provider string, cloudConfig map[string]string) error

	// simple operations
	Exists(path string) (bool, error)
	MakeDirectory(path string) error
	ListDirectory(path string) ([]string, error)

	// download and upload files/directories
	DownloadFile(path string, destination string) error
	UploadFile(path string, source string) error
	DownloadDirectory(path string, destination string) error
	UploadDirectory(path string, source string) error

	// delete files/directories
	DeleteFile(path string) error
	DeleteDirectory(path string) error
}
