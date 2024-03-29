package rq

import (
	"net/url"
	"reflect"
	"strconv"
)

// string型のURLクエリをセットする。
func Query[String ~string](key string, value String) Option {
	return OptionFunc(func(r *Request) {
		r.query.Set(key, string(value))
	})
}

// int型のURLクエリをセットする。
func QueryInt[Int ~int | ~int8 | ~int16 | ~int32 | ~int64](key string, value Int) Option {
	return Query(key, strconv.FormatInt(int64(value), 10))
}

// uint型のURLクエリをセットする。
func QueryUint[Uint ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr](key string, value Uint) Option {
	return Query(key, strconv.FormatUint(uint64(value), 10))
}

// bool型のURLクエリをセットする。
func QueryBool[Bool ~bool](key string, value Bool) Option {
	return Query(key, strconv.FormatBool(bool(value)))
}

// float型のURLクエリをセットする。
func QueryFloat[Float ~float32 | ~float64](key string, value Float) Option {
	return Query(key, strconv.FormatFloat(float64(value), 'f', -1, reflect.TypeOf(value).Bits()))
}

// url.Values型のURLクエリをセットする。
func QueryValues(values url.Values) Option {
	return OptionFunc(func(r *Request) {
		for key, value := range values {
			r.query[key] = value
		}
	})
}
