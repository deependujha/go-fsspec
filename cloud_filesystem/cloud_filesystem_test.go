package cloud_filesystem

import (
	"testing"
)

func TestCloudFileSystemIsValid(t *testing.T) {
	// cloud file system
	var _ AbstractCloudFileStorage = &S3FileStorage{} // means, it implements the interface

}
