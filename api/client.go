package api

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type FilebinClient struct {
	client  *resty.Client
	baseURL string
}

func NewClient() *FilebinClient {
	baseURL := "https://filebin.net"

	client := resty.New().
		SetBaseURL(baseURL).
		SetRetryCount(3).
		SetRetryWaitTime(time.Second).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "filebin-cli/2.0")

	return &FilebinClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *FilebinClient) EnableDebug() {
	c.client.SetDebug(true)
}

func (c *FilebinClient) GetClient() *resty.Client {
	return c.client
}
