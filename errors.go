package rq

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrBadRequest                    = statusCodeError(http.StatusBadRequest)
	ErrUnauthorized                  = statusCodeError(http.StatusUnauthorized)
	ErrPaymentRequired               = statusCodeError(http.StatusPaymentRequired)
	ErrForbidden                     = statusCodeError(http.StatusForbidden)
	ErrNotFound                      = statusCodeError(http.StatusNotFound)
	ErrMethodNotAllowed              = statusCodeError(http.StatusMethodNotAllowed)
	ErrNotAcceptable                 = statusCodeError(http.StatusNotAcceptable)
	ErrProxyAuthRequired             = statusCodeError(http.StatusProxyAuthRequired)
	ErrRequestTimeout                = statusCodeError(http.StatusRequestTimeout)
	ErrConflict                      = statusCodeError(http.StatusConflict)
	ErrGone                          = statusCodeError(http.StatusGone)
	ErrLengthRequired                = statusCodeError(http.StatusLengthRequired)
	ErrPreconditionFailed            = statusCodeError(http.StatusPreconditionFailed)
	ErrRequestEntityTooLarge         = statusCodeError(http.StatusRequestEntityTooLarge)
	ErrRequestURITooLong             = statusCodeError(http.StatusRequestURITooLong)
	ErrUnsupportedMediaType          = statusCodeError(http.StatusUnsupportedMediaType)
	ErrRequestedRangeNotSatisfiable  = statusCodeError(http.StatusRequestedRangeNotSatisfiable)
	ErrExpectationFailed             = statusCodeError(http.StatusExpectationFailed)
	ErrTeapot                        = statusCodeError(http.StatusTeapot)
	ErrMisdirectedRequest            = statusCodeError(http.StatusMisdirectedRequest)
	ErrUnprocessableEntity           = statusCodeError(http.StatusUnprocessableEntity)
	ErrLocked                        = statusCodeError(http.StatusLocked)
	ErrFailedDependency              = statusCodeError(http.StatusFailedDependency)
	ErrTooEarly                      = statusCodeError(http.StatusTooEarly)
	ErrUpgradeRequired               = statusCodeError(http.StatusUpgradeRequired)
	ErrPreconditionRequired          = statusCodeError(http.StatusPreconditionRequired)
	ErrTooManyRequests               = statusCodeError(http.StatusTooManyRequests)
	ErrRequestHeaderFieldsTooLarge   = statusCodeError(http.StatusRequestHeaderFieldsTooLarge)
	ErrUnavailableForLegalReasons    = statusCodeError(http.StatusUnavailableForLegalReasons)
	ErrInternalServerError           = statusCodeError(http.StatusInternalServerError)
	ErrNotImplemented                = statusCodeError(http.StatusNotImplemented)
	ErrBadGateway                    = statusCodeError(http.StatusBadGateway)
	ErrServiceUnavailable            = statusCodeError(http.StatusServiceUnavailable)
	ErrGatewayTimeout                = statusCodeError(http.StatusGatewayTimeout)
	ErrHTTPVersionNotSupported       = statusCodeError(http.StatusHTTPVersionNotSupported)
	ErrVariantAlsoNegotiates         = statusCodeError(http.StatusVariantAlsoNegotiates)
	ErrInsufficientStorage           = statusCodeError(http.StatusInsufficientStorage)
	ErrLoopDetected                  = statusCodeError(http.StatusLoopDetected)
	ErrNotExtended                   = statusCodeError(http.StatusNotExtended)
	ErrNetworkAuthenticationRequired = statusCodeError(http.StatusNetworkAuthenticationRequired)
)

func statusCodeError(statusCode int) error {
	return &StatusCodeError{StatusCode: statusCode}
}

type StatusCodeError struct {
	StatusCode int
}

func (err *StatusCodeError) Error() string {
	return fmt.Sprintf("invalid status code: %d %s", err.StatusCode, http.StatusText(err.StatusCode))
}

func (err *StatusCodeError) Is(target error) bool {
	if x, ok := target.(*StatusCodeError); ok {
		return err.StatusCode == x.StatusCode
	}
	return false
}

// ResponseError発生時に読み取るBodyの最大サイズをセットする。
func ErrBodyLimit(limit int64) Option {
	return OptionFunc(func(r *Request) {
		r.errBodyLimit = limit
	})
}

func responseError(response *http.Response, limit int64, err error) error {
	body, _ := io.ReadAll(io.LimitReader(response.Body, limit))
	return &ResponseError{
		Method:     response.Request.Method,
		URL:        response.Request.URL,
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       body,
		err:        err,
	}
}

type ResponseError struct {
	Method     string
	URL        *url.URL
	StatusCode int
	Header     http.Header
	Body       []byte
	err        error
}

func (err *ResponseError) Error() string {
	method := err.Method
	if method == "" {
		method = "Get"
	} else {
		method = err.Method[:1] + strings.ToLower(err.Method)[1:]
	}

	url := err.URL.String()
	if _, ok := err.URL.User.Password(); ok {
		url = strings.Replace(url, err.URL.User.String()+"@", err.URL.User.Username()+":***@", 1)
	}

	return fmt.Sprintf("%s %q: %s", method, url, err.err.Error())
}

func (err *ResponseError) Unwrap() error {
	return err.err
}

func AsResponseError(err error) (*ResponseError, bool) {
	e, ok := err.(*ResponseError)
	return e, ok
}

func MapResponseError(err error, fn func(*ResponseError) error) error {
	e, ok := err.(*ResponseError)
	if ok {
		return fn(e)
	}
	return e
}
