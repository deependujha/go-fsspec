// Package go_fsspec provides a unified interface for interacting with different cloud storage systems
// such as AWS S3, Google Cloud Storage, and Azure Blob Storage.
//
// go_fsspec offers a simple and powerful way to manage files and directories across multiple cloud providers
// without having to deal with provider-specific APIs or implementations.
//
// go_fsspec contains the following key functionalities:
//
// The core interface `CloudStorage` provides a set of common methods for file and directory operations,
// allowing users to interact with any supported cloud provider using the same interface.
//
// Supported operations include creating directories, listing files, downloading and uploading files,
// and deleting files and directories. Additional methods for downloading/uploading entire directories
// are also available for bulk operations.
//
// Each cloud provider (AWS S3, Google Cloud Storage, Azure Blob Storage) has its own implementation of the `CloudStorage` interface,
// allowing developers to switch providers with minimal code changes.
//
// Future versions of go_fsspec will include support for more cloud providers, advanced file versioning, and cross-cloud
// transfer capabilities.
//
// Contributions and feedback are welcome! Feel free to check out the repository and get involved.
//
// [AWS S3]: https://aws.amazon.com/s3/
// [Google Cloud Storage]: https://cloud.google.com/storage
// [Azure Blob Storage]: https://azure.microsoft.com/en-us/services/storage/blobs/
package go_fsspec
