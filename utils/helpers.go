package helpers

import (
	"errors"

	"github.com/deependujha/go-fsspec/cloud_filesystem"
	"github.com/deependujha/go-fsspec/constants"
)

// CheckIfProviderIsSupported checks if the given provider is supported
func CheckIfProviderIsSupported(provider string) bool {
	isProviderSupported := false

	for _, supportedProvider := range constants.SUPPORTED_CLOUD_PROVIDERS {
		if supportedProvider == provider {
			isProviderSupported = true
			break
		}
	}

	return isProviderSupported
}

// GetCloudFileStorage returns the CloudFileStorage implementation for the given provider
func GetCloudFileStorage(provider string, cloudConfig map[string]string) (cloud_filesystem.AbstractCloudFileStorage, error) {
	isProviderSupported := CheckIfProviderIsSupported(provider)

	if !isProviderSupported {
		return nil, errors.New("provider not supported")
	}

	var cloud_file_system cloud_filesystem.AbstractCloudFileStorage

	switch provider {
	case constants.S3:
		// cloud_file_system = &S3FileStorage{}
		return nil, errors.New("S3 not implemented")
	case constants.GCS:
		// cloud_file_system = &GCSFileStorage{}
		return nil, errors.New("GCS not implemented")
	case constants.ABFS:
		// cloud_file_system = &ABFSFileStorage{}
		return nil, errors.New("ABFS not implemented")
	}

	return cloud_file_system, nil
}
