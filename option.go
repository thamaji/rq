package rq

import (
	"context"
	"net/http"
)

type Option interface {
	Apply(*Request)
}

type OptionFunc func(*Request)

func (f OptionFunc) Apply(r *Request) {
	f(r)
}

// コンテキストをセットする。
func Context(ctx context.Context) Option {
	return OptionFunc(func(r *Request) {
		r.ctx = ctx
	})
}

// HTTPクライアントをセットする。
func HTTPClient(client *http.Client) Option {
	return OptionFunc(func(r *Request) {
		r.client = client
	})
}
