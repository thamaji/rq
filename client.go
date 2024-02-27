package rq

import (
	"net/http"
)

// クライアントをつくる。
func NewClient(options ...Option) *Client {
	return &Client{options: options}
}

type Client struct {
	options []Option
}

// リクエストをつくる。
func (c Client) NewRequest(method string, url string, options ...Option) *Request {
	return NewRequest(method, url, c.options...).With(options...)
}

// GETメソッドのリクエストをつくる。
func (c Client) Get(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodGet, url, options...)
}

// HEADメソッドのリクエストをつくる。
func (c Client) Head(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodHead, url, options...)
}

// POSTメソッドのリクエストをつくる。
func (c Client) Post(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodPost, url, options...)
}

// PUTメソッドのリクエストをつくる。
func (c Client) Put(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodPut, url, options...)
}

// PATCHメソッドのリクエストをつくる。
func (c Client) Patch(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodPatch, url, options...)
}

// DELETEメソッドのリクエストをつくる。
func (c Client) Delete(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodDelete, url, options...)
}

// CONNECTメソッドのリクエストをつくる。
func (c Client) Connect(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodConnect, url, options...)
}

// OPTIONSメソッドのリクエストをつくる。
func (c Client) Options(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodOptions, url, options...)
}

// TRACEメソッドのリクエストをつくる。
func (c Client) Trace(url string, options ...Option) *Request {
	return c.NewRequest(http.MethodTrace, url, options...)
}
