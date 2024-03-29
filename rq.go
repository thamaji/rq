package rq

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	_url "net/url"
)

// GETメソッドのリクエストをつくる。
func Get(url string, options ...Option) *Request {
	return NewRequest(http.MethodGet, url, options...)
}

// HEADメソッドのリクエストをつくる。
func Head(url string, options ...Option) *Request {
	return NewRequest(http.MethodHead, url, options...)
}

// POSTメソッドのリクエストをつくる。
func Post(url string, options ...Option) *Request {
	return NewRequest(http.MethodPost, url, options...)
}

// PUTメソッドのリクエストをつくる。
func Put(url string, options ...Option) *Request {
	return NewRequest(http.MethodPut, url, options...)
}

// PATCHメソッドのリクエストをつくる。
func Patch(url string, options ...Option) *Request {
	return NewRequest(http.MethodPatch, url, options...)
}

// DELETEメソッドのリクエストをつくる。
func Delete(url string, options ...Option) *Request {
	return NewRequest(http.MethodDelete, url, options...)
}

// CONNECTメソッドのリクエストをつくる。
func Connect(url string, options ...Option) *Request {
	return NewRequest(http.MethodConnect, url, options...)
}

// OPTIONSメソッドのリクエストをつくる。
func Options(url string, options ...Option) *Request {
	return NewRequest(http.MethodOptions, url, options...)
}

// TRACEメソッドのリクエストをつくる。
func Trace(url string, options ...Option) *Request {
	return NewRequest(http.MethodTrace, url, options...)
}

// リクエストをつくる。
func NewRequest(method string, url string, options ...Option) *Request {
	r := &Request{
		method:       method,
		url:          url,
		body:         nil,
		err:          nil,
		errBodyLimit: 4096,
		query:        _url.Values{},
		header:       http.Header{},
		client:       http.DefaultClient,
	}
	return r.With(options...)
}

// リクエスト。
type Request struct {
	method string
	url    string
	body   io.Reader
	err    error

	errBodyLimit int64

	preHook  []func(*http.Request) error
	postHook []func(*http.Response) error

	query  _url.Values
	header http.Header
	client *http.Client
	ctx    context.Context
}

// リクエストにオプションを適用する。
func (r *Request) With(options ...Option) *Request {
	for _, option := range options {
		option.Apply(r)
	}
	return r
}

// リクエストを実行してレスポンスを返す。
func (r *Request) Do() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	url, err := _url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	for k, v := range r.query {
		query[k] = v
	}
	url.RawQuery = query.Encode()

	request, err := http.NewRequest(r.method, url.String(), r.body)
	if err != nil {
		return nil, err
	}

	if r.ctx != nil {
		request.WithContext(r.ctx)
	}

	if request.Header == nil {
		request.Header = http.Header{}
	}
	for k, v := range r.header {
		request.Header[k] = v
	}

	for _, hook := range r.preHook {
		if err := hook(request); err != nil {
			return nil, err
		}
	}

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	body := response.Body
	response.Body = readCloser{
		read: body.Read,
		close: func() error {
			_, err := io.Copy(io.Discard, body)
			if err1 := body.Close(); err == nil {
				err = err1
			}
			return err
		},
	}

	for _, hook := range r.postHook {
		if err := hook(response); err != nil {
			response.Body.Close()
			return nil, err
		}
	}

	return response, nil
}

// リクエストを実行してレスポンスボディをひらく。
func (r *Request) Open() (io.ReadCloser, error) {
	response, err := r.Do()
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		err := responseError(response, r.errBodyLimit, statusCodeError(response.StatusCode))
		response.Body.Close()
		return nil, err
	}

	return response.Body, nil
}

// リクエストを実行する。
func (r *Request) Done() error {
	body, err := r.Open()
	if err != nil {
		return err
	}
	return body.Close()
}

// リクエストを実行してレスポンスボディのバイト列を返す。
func (r *Request) Fetch() ([]byte, error) {
	body, err := r.Open()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, body)
	if err1 := body.Close(); err == nil {
		err = err1
	}
	return buf.Bytes(), err
}

// リクエストを実行してレスポンスボディのJSONをパースする。
func (r *Request) FetchJSON(v any) error {
	r.With(Accept("application/json"))
	body, err := r.Open()
	if err != nil {
		return err
	}
	err = json.NewDecoder(body).Decode(v)
	if err1 := body.Close(); err == nil {
		err = err1
	}
	return err
}

type readCloser struct {
	read  func([]byte) (int, error)
	close func() error
}

func (rc readCloser) Read(p []byte) (int, error) {
	return rc.read(p)
}
func (rc readCloser) Close() error {
	return rc.close()
}
