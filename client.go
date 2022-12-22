package rq

import "net/http"

// HTTPクライアントをセットする。
func Client(client *http.Client) Option {
	return OptionFunc(func(r *Request) {
		r.client = client
	})
}
