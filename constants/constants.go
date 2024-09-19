package constants

const (
	S3   string = "s3"
	GCS  string = "gcs"
	ABFS string = "abfs"
)

var SUPPORTED_CLOUD_PROVIDERS = [...]string{S3, GCS, ABFS}
