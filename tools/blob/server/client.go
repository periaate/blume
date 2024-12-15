package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/periaate/blob"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes a new API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// Get retrieves a blob from the server
func (c *Client) Get(bucket, name string) (r io.ReadCloser, ct blob.ContentType, err error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, bucket, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		err = fmt.Errorf("GET failed: %s", resp.Status)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	ct = blob.GetCT(contentType)
	r = resp.Body

	return
}

// Set uploads a blob to the server
func (c *Client) Set(bucket, name string, data io.Reader, ct blob.ContentType) (err error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, bucket, name)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", ct.String())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST failed: %s", resp.Status)
	}

	return
}

// Delete removes a blob from the server
func (c *Client) Delete(bucket, name string) (err error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, bucket, name)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("DELETE failed: %s", resp.Status)
	}

	return
}

// ValidateBaseURL ensures the base URL is valid and normalized
func (c *Client) ValidateBaseURL() (err error) {
	parsed, err := url.Parse(c.BaseURL)
	if err != nil || !parsed.IsAbs() {
		return errors.New("invalid base URL")
	}
	c.BaseURL = parsed.String()
	return
}
