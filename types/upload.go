package types

import (
	"path/filepath"
)

type UploadConfig struct {
	BinID    string
	FilePath string
	Filename string
	ClientID string
}

func NewUploadConfig(binID, filePath string) UploadConfig {
	filename := filepath.Base(filePath)
	return UploadConfig{
		BinID:    binID,
		FilePath: filePath,
		Filename: filename,
	}
}

func (u *UploadConfig) SetCustomFilename(filename string) {
	u.Filename = filename
}

func (u *UploadConfig) SetClientID(clientID string) {
	u.ClientID = clientID
}
