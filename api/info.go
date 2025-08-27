package api

import (
	"fmt"

	"filebin-cli/types"
)

func (c *FilebinClient) GetBinInfo(binID string) (*types.BinInfo, error) {
	if binID == "" {
		return nil, fmt.Errorf("bin ID cannot be empty")
	}

	var result types.BinInfo
	resp, err := c.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("/%s", binID))

	if err != nil {
		return nil, fmt.Errorf("failed to fetch bin info: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("bin not found: %s", binID)
	}

	return &result, nil
}
