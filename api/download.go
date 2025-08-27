package api

import (
	"fmt"

	"filebin-cli/errors"
	"filebin-cli/types"
)

func (c *FilebinClient) DownloadResource(config types.DownloadConfig) error {
	var url string
	if config.Resource.IsFile() {
		url = config.Resource.URL()
	} else {
		url = config.Resource.ArchiveURL()
	}

	resp, err := c.client.R().
		SetHeader("User-Agent", "curl/wget/VLC").
		SetOutput(config.GetOutputPath()).
		Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	case 404:
		if config.Resource.IsFile() {
			return &errors.FileNotFoundError{
				BinID:    config.Resource.BinID,
				Filename: config.Resource.Filename,
			}
		}
		return &errors.BinNotFoundError{BinID: config.Resource.BinID}
	default:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "download failed",
		}
	}
}
