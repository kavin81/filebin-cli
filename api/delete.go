package api

import (
	"fmt"

	"filebin-cli/errors"
	"filebin-cli/types"
)

func (c *FilebinClient) DeleteResource(resource types.Resource) error {
	if resource.BinID == "" {
		return fmt.Errorf("bin ID cannot be empty")
	}

	resp, err := c.client.R().Delete(resource.URL())
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	case 404:
		if resource.IsFile() {
			return &errors.FileNotFoundError{BinID: resource.BinID, Filename: resource.Filename}
		}
		return &errors.BinNotFoundError{BinID: resource.BinID}
	default:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "failed to delete resource",
		}
	}
}
