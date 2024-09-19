# go-fsspec

![go fsspec](./go-fsspec.png)

---

## Idea

- A simple yet powerful, library to manage files and directories in cloud storage systems (`AWS S3, Google Cloud Storage, Azure Blob Storage,` etc.).

---

## Features

- We will have an **`abstractCloudStorage`** interface that will be implemented by each storage system.
- `abstractCloudStorage` will have these methods:
  - `init(cloudProvider string, cloudConfig map[string]string) error`
  - `fileExists(path string) (bool, error)`
  - `makeDirectory(path string) error`
  - `listDirectory(path string) ([]string, error)` / `ls(path string) ([]string, error)`
  - `downloadFile(path string, destination string) error`
  - `uploadFile(path string, source string) error`
  - `downloadDirectory(path string, destination string) error`
  - `uploadDirectory(path string, source string) error`
  - `deleteFile(path string) error`
  - `deleteDirectory(path string) error`

- Each storage system will have its own implementation of `abstractCloudStorage` interface.
