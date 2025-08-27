package types

import "path/filepath"

// DownloadConfig represents the configuration for a download operation.
type DownloadConfig struct {
	OutputPath string
	Resource   Resource
}

// GetOutputPath returns the output path for the download operation.
func (dc DownloadConfig) GetOutputPath() string {
	if dc.Resource.IsFile() {
		return filepath.Join(dc.OutputPath, dc.Resource.Filename)
	}
	return filepath.Join(dc.OutputPath, dc.Resource.BinID+".tar")
}
