package api

import (
	"fmt"
	"time"

	"filebin-cli/errors"
	"filebin-cli/types"
)

type Operations struct {
	client *FilebinClient
}

func NewOperations() *Operations {
	return &Operations{
		client: NewClient(),
	}
}

func (ops *Operations) EnableDebug() {
	ops.client.EnableDebug()
}

func (ops *Operations) GetBinInfo(binID string) (*types.BinInfo, error) {
	return ops.client.GetBinInfo(binID)
}

func (ops *Operations) DeleteResource(resource types.Resource) error {
	return ops.client.DeleteResource(resource)
}

func (ops *Operations) LockBin(binID string) error {
	return ops.client.LockBin(binID)
}

func (ops *Operations) DownloadResource(config types.DownloadConfig) error {
	return ops.client.DownloadResource(config)
}

func (ops *Operations) UploadFile(config types.UploadConfig) error {
	return ops.client.UploadFile(config)
}

func (ops *Operations) SetTimeout(timeout time.Duration) {
	ops.client.GetClient().SetTimeout(timeout)
}

func (ops *Operations) CheckHealth() error {
	resp, err := ops.client.GetClient().R().Get("/health")
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}

	if resp.StatusCode() != 200 {
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    "service unavailable",
		}
	}

	return nil
}
