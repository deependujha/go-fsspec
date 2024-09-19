package go_fsspec

import (
	"errors"

	"github.com/deependujha/go-fsspec/cloud_filesystem"
	helpers "github.com/deependujha/go-fsspec/utils"
)

// CloudFileStorage is a struct that implements AbstractCloudFileStorage interface
type CloudFileStorage struct {
	provider    string
	cloudConfig map[string]string
	fs          cloud_filesystem.AbstractCloudFileStorage
}

// init initializes the CloudFileStorage struct
func (cfs *CloudFileStorage) init(provider string, cloudConfig map[string]string) error {
	isProviderSupported := helpers.CheckIfProviderIsSupported(provider)

	if !isProviderSupported {
		return errors.New("provider not supported")
	}

	fs, err := helpers.GetCloudFileStorage(provider, cloudConfig)

	if err != nil {
		return err
	}

	cfs.provider = provider
	cfs.cloudConfig = cloudConfig
	cfs.fs = fs

	return nil
}

// check if the file/directory exists
func (cfs *CloudFileStorage) Exists(path string) (bool, error) {
	if cfs.fs == nil {
		return false, errors.New("cloud file storage not initialized")
	}

	return cfs.fs.Exists(path)
}

// make directory in the cloud storage
func (cfs *CloudFileStorage) MakeDirectory(path string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.MakeDirectory(path)
}

// list directory contents
func (cfs *CloudFileStorage) ListDirectory(path string) ([]string, error) {
	if cfs.fs == nil {
		return nil, errors.New("cloud file storage not initialized")
	}

	return cfs.fs.ListDirectory(path)
}

// download file from the cloud storage
func (cfs *CloudFileStorage) DownloadFile(path string, destination string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.DownloadFile(path, destination)
}

// upload file to the cloud storage
func (cfs *CloudFileStorage) UploadFile(path string, source string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.UploadFile(path, source)
}

// download directory from the cloud storage
func (cfs *CloudFileStorage) DownloadDirectory(path string, destination string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.DownloadDirectory(path, destination)
}

// upload directory to the cloud storage
func (cfs *CloudFileStorage) UploadDirectory(path string, source string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.UploadDirectory(path, source)
}

// delete file from the cloud storage
func (cfs *CloudFileStorage) DeleteFile(path string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.DeleteFile(path)
}

// delete directory from the cloud storage
func (cfs *CloudFileStorage) DeleteDirectory(path string) error {
	if cfs.fs == nil {
		return errors.New("cloud file storage not initialized")
	}

	return cfs.fs.DeleteDirectory(path)
}
