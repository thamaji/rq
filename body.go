package rq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

// リクエストボディをセットする。
func Body(body io.Reader) Option {
	return OptionFunc(func(r *Request) {
		r.body = body
	})
}

// バイト列のリクエストボディをセットする。
func BodyBytes(body []byte) Option {
	return OptionFunc(func(r *Request) {
		r.With(
			ContentLength(len(body)),
			Body(bytes.NewReader(body)),
		)
	})
}

// 文字列のリクエストボディをセットする。
func BodyString(body string) Option {
	return BodyBytes([]byte(body))
}

// JSONのリクエストボディをセットする。
func BodyJSON[V any](body V) Option {
	return OptionFunc(func(r *Request) {
		b, err := json.Marshal(body)
		if err != nil {
			r.err = err
			return
		}

		r.With(
			ContentType("application/json").Charset("UTF-8"),
			BodyBytes(b),
		)
	})
}

// x-www-form-urlencodedのリクエストボディをセットする。
func BodyFormURLEncoded(form url.Values) Option {
	return OptionFunc(func(r *Request) {
		r.With(
			ContentType("application/x-www-form-urlencoded").Charset("UTF-8"),
			BodyString(form.Encode()),
		)
	})
}

// 関数の実行結果をリクエストボディとしてセットする。
func BodyFunc(f func() (io.Reader, error)) Option {
	return OptionFunc(func(r *Request) {
		r.body, r.err = f()
	})
}
