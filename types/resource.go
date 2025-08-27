package types

import "fmt"

// Resource represents a bin or a file within a bin.
type Resource struct {
	BinID    string
	Filename string
}

// NewBinResource creates a new Resource for a bin.
func NewBinResource(binID string) Resource {
	return Resource{BinID: binID}
}

// NewFileResource creates a new Resource for a file within a bin.
func NewFileResource(binID, filename string) Resource {
	return Resource{BinID: binID, Filename: filename}
}

// IsFile checks if the resource is a file.
func (r Resource) IsFile() bool {
	return r.Filename != ""
}

// URL returns the URL for the resource.
func (r Resource) URL() string {
	if r.IsFile() {
		return fmt.Sprintf("/%s/%s", r.BinID, r.Filename)
	}
	return fmt.Sprintf("/%s", r.BinID)
}

// ArchiveURL returns the archive URL for the resource.
func (r Resource) ArchiveURL() string {
	return fmt.Sprintf("/archive/%s/tar", r.BinID)
}
