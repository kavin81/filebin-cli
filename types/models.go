// Package types defines the data structures used in the filebin application.
package types

import "time"

// BinInfo represents information about a bin.
type BinInfo struct {
	Bin   BinMetadata `json:"bin"`
	Files []FileInfo  `json:"files"`
}

// BinMetadata represents metadata about a bin.
type BinMetadata struct {
	ID                string    `json:"id"`
	Readonly          bool      `json:"readonly"`
	Bytes             int64     `json:"bytes"`
	BytesReadable     string    `json:"bytes_readable"`
	Files             int       `json:"files"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedAtRelative string    `json:"updated_at_relative"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedAtRelative string    `json:"created_at_relative"`
	ExpiredAt         time.Time `json:"expired_at"`
	ExpiredAtRelative string    `json:"expired_at_relative"`
}

// FileInfo represents information about a file within a bin.
type FileInfo struct {
	Filename          string    `json:"filename"`
	ContentType       string    `json:"content-type"`
	Bytes             int64     `json:"bytes"`
	BytesReadable     string    `json:"bytes_readable"`
	MD5               string    `json:"md5"`
	SHA256            string    `json:"sha256"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedAtRelative string    `json:"updated_at_relative"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedAtRelative string    `json:"created_at_relative"`
}
