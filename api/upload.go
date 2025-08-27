package api

import (
	"fmt"
	"os"
	"strings"

	"filebin-cli/errors"
	"filebin-cli/types"
)

func (c *FilebinClient) UploadFile(config types.UploadConfig) error {
	// Open the file
	file, err := os.Open(config.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file info for size
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	url := fmt.Sprintf("/%s/%s", config.BinID, config.Filename)

	req := c.client.R().
		SetFileReader("file", config.Filename, file).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Add custom client ID if provided
	if config.ClientID != "" {
		req.SetHeader("cid", config.ClientID)
	}

	resp, err := req.Post(url)
	if err != nil {
		return fmt.Errorf("upload failed: %w", err)
	}

	switch resp.StatusCode() {
	case 201:
		return nil
	case 400:
		// Check for specific error messages in the response body
		responseBody := string(resp.Body())
		if strings.Contains(responseBody, "The bin is too short") {
			return &errors.APIError{
				StatusCode: resp.StatusCode(),
				Message:    "the bin ID is too short",
			}
		}
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "invalid input, typically invalid bin or filename specified",
		}
	case 403:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "storage limitation was reached",
		}
	case 405:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "the bin is locked and can not be written to",
		}
	default:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "upload failed",
		}
	}
}
