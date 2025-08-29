package api

import (
	"fmt"
	"os"
	"strings"

	"filebin-cli/errors"
	"filebin-cli/types"
)

func (c *FilebinClient) UploadFile(config types.UploadConfig) error {
	fileData, err := os.ReadFile(config.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	url := fmt.Sprintf("/%s/%s", config.BinID, config.Filename)

	req := c.client.R().
		SetBody(fileData).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Accept", "application/json")

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
