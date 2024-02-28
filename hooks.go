package rq

import "net/http"

// リクエスト実行直前のフックをセットする。
func PreHook(hook func(*http.Request) error) Option {
	return OptionFunc(func(r *Request) {
		r.preHook = append(r.preHook, hook)
	})
}

// リクエスト実行直後のフックをセットする。
func PostHook(hook func(*http.Response) error) Option {
	return OptionFunc(func(r *Request) {
		r.postHook = append(r.postHook, hook)
	})
}
