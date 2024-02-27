package rq

import (
	"path"
	"strings"
)

// URLを結合する。
func URL(url string, elem ...string) string {
	if len(elem) <= 0 {
		return url
	}
	return strings.TrimSuffix(url, "/") + "/" + strings.TrimPrefix(path.Join(elem...), "/")
}

// ベースURLをセットする。
func BaseURL(url string) Option {
	return OptionFunc(func(r *Request) {
		r.url = URL(url, r.url)
	})
}
