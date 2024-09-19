package cloud_filesystem

type S3FileStorage struct {
	// TODO: implement
}

func (s3fs *S3FileStorage) Init(provider string, cloudConfig map[string]string) error {
	// TODO: implement
	return nil
}

func (s3fs *S3FileStorage) Exists(path string) (bool, error) {
	// TODO: implement
	return false, nil
}

func (s3fs *S3FileStorage) MakeDirectory(path string) error {
	// TODO: implement
	return nil
}
func (s3fs *S3FileStorage) ListDirectory(path string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (s3fs *S3FileStorage) DownloadFile(path string, destination string) error {
	// TODO: implement
	return nil
}
func (s3fs *S3FileStorage) UploadFile(path string, source string) error {
	// TODO: implement
	return nil
}
func (s3fs *S3FileStorage) DownloadDirectory(path string, destination string) error {
	// TODO: implement
	return nil
}
func (s3fs *S3FileStorage) UploadDirectory(path string, source string) error {
	// TODO: implement
	return nil
}

// deleteFile deletes a file from the cloud storage
func (s3fs *S3FileStorage) DeleteFile(path string) error {
	// TODO: implement
	return nil
}

// deleteDirectory deletes a directory from the cloud storage
func (s3fs *S3FileStorage) DeleteDirectory(path string) error {
	// TODO: implement
	return nil
}
