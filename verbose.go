package rq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// リクエスト、レスポンスの内容を出力する。
func Verbose(w io.Writer) Option {
	return OptionFunc(func(r *Request) {
		r.With(
			PreHook(func(request *http.Request) error {
				fmt.Fprintln(w, ">", request.Method, request.URL.Path, request.Proto)
				fmt.Fprintln(w, ">", "Host:", request.Host)
				for k := range request.Header {
					fmt.Fprintln(w, ">", k+":", request.Header.Get(k))
				}
				return nil
			}),
			PostHook(func(response *http.Response) error {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}
				response.Body = io.NopCloser(bytes.NewReader(body))

				fmt.Fprintln(w, "<", response.Proto, response.StatusCode)
				for k := range response.Header {
					fmt.Fprintln(w, "<", k+":", response.Header.Get(k))
				}
				fmt.Fprintln(w, "<")
				fmt.Fprintln(w, string(body))
				return nil
			}),
		)
	})
}
