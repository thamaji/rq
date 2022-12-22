package rq

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest                    = errors.New(http.StatusText(http.StatusBadRequest))
	ErrUnauthorized                  = errors.New(http.StatusText(http.StatusUnauthorized))
	ErrPaymentRequired               = errors.New(http.StatusText(http.StatusPaymentRequired))
	ErrForbidden                     = errors.New(http.StatusText(http.StatusForbidden))
	ErrNotFound                      = errors.New(http.StatusText(http.StatusNotFound))
	ErrMethodNotAllowed              = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	ErrNotAcceptable                 = errors.New(http.StatusText(http.StatusNotAcceptable))
	ErrProxyAuthRequired             = errors.New(http.StatusText(http.StatusProxyAuthRequired))
	ErrRequestTimeout                = errors.New(http.StatusText(http.StatusRequestTimeout))
	ErrConflict                      = errors.New(http.StatusText(http.StatusConflict))
	ErrGone                          = errors.New(http.StatusText(http.StatusGone))
	ErrLengthRequired                = errors.New(http.StatusText(http.StatusLengthRequired))
	ErrPreconditionFailed            = errors.New(http.StatusText(http.StatusPreconditionFailed))
	ErrRequestEntityTooLarge         = errors.New(http.StatusText(http.StatusRequestEntityTooLarge))
	ErrRequestURITooLong             = errors.New(http.StatusText(http.StatusRequestURITooLong))
	ErrUnsupportedMediaType          = errors.New(http.StatusText(http.StatusUnsupportedMediaType))
	ErrRequestedRangeNotSatisfiable  = errors.New(http.StatusText(http.StatusRequestedRangeNotSatisfiable))
	ErrExpectationFailed             = errors.New(http.StatusText(http.StatusExpectationFailed))
	ErrTeapot                        = errors.New(http.StatusText(http.StatusTeapot))
	ErrMisdirectedRequest            = errors.New(http.StatusText(http.StatusMisdirectedRequest))
	ErrUnprocessableEntity           = errors.New(http.StatusText(http.StatusUnprocessableEntity))
	ErrLocked                        = errors.New(http.StatusText(http.StatusLocked))
	ErrFailedDependency              = errors.New(http.StatusText(http.StatusFailedDependency))
	ErrTooEarly                      = errors.New(http.StatusText(http.StatusTooEarly))
	ErrUpgradeRequired               = errors.New(http.StatusText(http.StatusUpgradeRequired))
	ErrPreconditionRequired          = errors.New(http.StatusText(http.StatusPreconditionRequired))
	ErrTooManyRequests               = errors.New(http.StatusText(http.StatusTooManyRequests))
	ErrRequestHeaderFieldsTooLarge   = errors.New(http.StatusText(http.StatusRequestHeaderFieldsTooLarge))
	ErrUnavailableForLegalReasons    = errors.New(http.StatusText(http.StatusUnavailableForLegalReasons))
	ErrInternalServerError           = errors.New(http.StatusText(http.StatusInternalServerError))
	ErrNotImplemented                = errors.New(http.StatusText(http.StatusNotImplemented))
	ErrBadGateway                    = errors.New(http.StatusText(http.StatusBadGateway))
	ErrServiceUnavailable            = errors.New(http.StatusText(http.StatusServiceUnavailable))
	ErrGatewayTimeout                = errors.New(http.StatusText(http.StatusGatewayTimeout))
	ErrHTTPVersionNotSupported       = errors.New(http.StatusText(http.StatusHTTPVersionNotSupported))
	ErrVariantAlsoNegotiates         = errors.New(http.StatusText(http.StatusVariantAlsoNegotiates))
	ErrInsufficientStorage           = errors.New(http.StatusText(http.StatusInsufficientStorage))
	ErrLoopDetected                  = errors.New(http.StatusText(http.StatusLoopDetected))
	ErrNotExtended                   = errors.New(http.StatusText(http.StatusNotExtended))
	ErrNetworkAuthenticationRequired = errors.New(http.StatusText(http.StatusNetworkAuthenticationRequired))
)

func StatusError(code int) error {
	switch code {
	default:
		return nil
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusPaymentRequired:
		return ErrPaymentRequired
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	case http.StatusNotAcceptable:
		return ErrNotAcceptable
	case http.StatusProxyAuthRequired:
		return ErrProxyAuthRequired
	case http.StatusRequestTimeout:
		return ErrRequestTimeout
	case http.StatusConflict:
		return ErrConflict
	case http.StatusGone:
		return ErrGone
	case http.StatusLengthRequired:
		return ErrLengthRequired
	case http.StatusPreconditionFailed:
		return ErrPreconditionFailed
	case http.StatusRequestEntityTooLarge:
		return ErrRequestEntityTooLarge
	case http.StatusRequestURITooLong:
		return ErrRequestURITooLong
	case http.StatusUnsupportedMediaType:
		return ErrUnsupportedMediaType
	case http.StatusRequestedRangeNotSatisfiable:
		return ErrRequestedRangeNotSatisfiable
	case http.StatusExpectationFailed:
		return ErrExpectationFailed
	case http.StatusTeapot:
		return ErrTeapot
	case http.StatusMisdirectedRequest:
		return ErrMisdirectedRequest
	case http.StatusUnprocessableEntity:
		return ErrUnprocessableEntity
	case http.StatusLocked:
		return ErrLocked
	case http.StatusFailedDependency:
		return ErrFailedDependency
	case http.StatusTooEarly:
		return ErrTooEarly
	case http.StatusUpgradeRequired:
		return ErrUpgradeRequired
	case http.StatusPreconditionRequired:
		return ErrPreconditionRequired
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusRequestHeaderFieldsTooLarge:
		return ErrRequestHeaderFieldsTooLarge
	case http.StatusUnavailableForLegalReasons:
		return ErrUnavailableForLegalReasons
	case http.StatusInternalServerError:
		return ErrInternalServerError
	case http.StatusNotImplemented:
		return ErrNotImplemented
	case http.StatusBadGateway:
		return ErrBadGateway
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	case http.StatusGatewayTimeout:
		return ErrGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return ErrHTTPVersionNotSupported
	case http.StatusVariantAlsoNegotiates:
		return ErrVariantAlsoNegotiates
	case http.StatusInsufficientStorage:
		return ErrInsufficientStorage
	case http.StatusLoopDetected:
		return ErrLoopDetected
	case http.StatusNotExtended:
		return ErrNotExtended
	case http.StatusNetworkAuthenticationRequired:
		return ErrNetworkAuthenticationRequired
	}
}
