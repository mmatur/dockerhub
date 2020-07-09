package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var closeFunc = func(body io.ReadCloser) {
	if body != nil {
		if errClose := body.Close(); errClose != nil {
			fmt.Print(errClose)
		}
	}
}

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// ClientOption allows to configure a Client.
// It is used in NewClient.
type ClientOption func(c *Client)

func (c *Client) doReq(ctx context.Context, method, url string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req.WithContext(ctx))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("docker: non-200 response code: could not read response body: %w", err)
		}

		return nil, fmt.Errorf("docker: non-200 response code: %s", string(b))
	}

	return resp.Body, nil
}

func (c *Client) ListAllRepo(ctx context.Context) (Results, error) {
	url := "https://hub.docker.com/api/content/v1/products/search?page_size=1000&q=&type=image"
	var r Results

	for url != "" {
		body, err := c.doReq(ctx, http.MethodGet, url, nil)
		defer closeFunc(body)
		if err != nil {
			return nil, err
		}

		var page Page
		if err = json.NewDecoder(body).Decode(&page); err != nil {
			return nil, err
		}
		r = append(r, page.Results...)
		url = page.Next
	}

	return r, nil
}
