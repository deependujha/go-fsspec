package go_fsspec

type AbstractCloudFileStorage interface {
	init(provider string, storage_options interface{})
}
