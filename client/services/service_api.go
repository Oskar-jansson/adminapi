package services

import (
	"context"
	"net/http"
	"net/url"
)

type clientInterface interface {
	DoRequest(ctx context.Context, method, endpoint string, params url.Values, body any, useAccessToken bool) (*http.Response, error)
	SetToken(v string)
	GetToken() string
}

type RequestBuilder struct {
	client    clientInterface
	ctx       context.Context
	method    string
	endpoint  string
	params    url.Values
	body      any
	useBearer bool
}

func NewRequest(ctx context.Context, c clientInterface) *RequestBuilder {
	return &RequestBuilder{
		client: c,
		ctx:    ctx,
		params: url.Values{},
	}
}

func (rb *RequestBuilder) Method(method, endpoint string) *RequestBuilder {
	rb.method = method
	rb.endpoint = endpoint
	return rb
}

func (rb *RequestBuilder) Query(params map[string]string) *RequestBuilder {
	for k, v := range params {
		rb.params.Add(k, v)
	}
	return rb
}

func (rb *RequestBuilder) QueryValues(values url.Values) *RequestBuilder {
	for k, vs := range values {
		for _, v := range vs {
			rb.params.Add(k, v)
		}
	}
	return rb
}

func (rb *RequestBuilder) Body(body any) *RequestBuilder {
	rb.body = body
	return rb
}

func (rb *RequestBuilder) WithBearer() *RequestBuilder {
	rb.useBearer = true
	return rb
}

func (rb *RequestBuilder) Do() (*http.Response, error) {
	return rb.client.DoRequest(rb.ctx, rb.method, rb.endpoint, rb.params, rb.body, rb.useBearer)
}
