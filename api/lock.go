package api

import (
	"fmt"

	"filebin-cli/errors"
)

func (c *FilebinClient) LockBin(binID string) error {
	if binID == "" {
		return fmt.Errorf("bin ID cannot be empty")
	}

	resp, err := c.client.R().Put(fmt.Sprintf("/%s", binID))
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	case 404:
		return &errors.BinNotFoundError{BinID: binID}
	default:
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "failed to lock bin",
		}
	}
}
