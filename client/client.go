package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
)

var ErrFailedParsingBody error = errors.New("failed parsing body")
var ErrFailedAssemblingRequest error = errors.New("failed assembling request")
var ErrFailedBuildingURL error = errors.New("failed building URL")

type Client struct {
	http    *http.Client
	baseUrl string
	token   string
	mu      sync.Mutex
}

func NewClient(httpClient http.Client, baseUrl string) *Client {
	return &Client{
		http:    &httpClient,
		baseUrl: baseUrl,
	}
}

func (c *Client) SetToken(v string) {
	c.mu.Lock()
	c.token = v
	c.mu.Unlock()
}

func (c *Client) GetToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.token
}

type RequestOption func(*http.Request)

func WithBearer(token string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("Access-Token", token)
	}
}

func toBody(v any) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}

	body, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Join(err, ErrFailedParsingBody)
	}

	return bytes.NewReader(body), nil
}

func (c *Client) DoRequest(ctx context.Context, method, endpoint string, params url.Values, body any, useAccessToken bool) (*http.Response, error) {
	fullUrl, err := c.buildURL(endpoint, params)
	if err != nil {
		return nil, errors.Join(err, ErrFailedAssemblingRequest)
	}

	bodyReader, err := toBody(body)
	if err != nil {
		return nil, errors.Join(err, ErrFailedAssemblingRequest)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullUrl, bodyReader)
	if err != nil {
		return nil, errors.Join(err, ErrFailedAssemblingRequest)
	}

	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if useAccessToken {
		req.Header.Set("Access-Token", c.GetToken())
	}

	return c.http.Do(req)
}

//
// Query params handeling
//

func (c *Client) buildURL(endpoint string, params url.Values) (string, error) {
	base, err := url.Parse(c.baseUrl)
	if err != nil {
		return "", errors.Join(err, ErrFailedBuildingURL)
	}

	endpoint = strings.TrimPrefix(endpoint, "/")
	base.Path = path.Join(base.Path, endpoint)
	base.RawQuery = params.Encode()

	return base.String(), nil
}
