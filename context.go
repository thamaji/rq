package rq

import "context"

// コンテキストをセットする。
func Context(ctx context.Context) Option {
	return OptionFunc(func(r *Request) {
		r.ctx = ctx
	})
}
