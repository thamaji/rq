package rq

import (
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

// string型のURLクエリをセットする。
func Query[String ~string](key string, value String) Option {
	return OptionFunc(func(r *Request) {
		r.query.Set(key, string(value))
	})
}

// int型のURLクエリをセットする。
func QueryInt[Int constraints.Signed](key string, value Int) Option {
	return Query(key, strconv.FormatInt(int64(value), 10))
}

// uint型のURLクエリをセットする。
func QueryUint[Uint constraints.Unsigned](key string, value Uint) Option {
	return Query(key, strconv.FormatUint(uint64(value), 10))
}

// bool型のURLクエリをセットする。
func QueryBool[Bool ~bool](key string, value Bool) Option {
	return Query(key, strconv.FormatBool(bool(value)))
}

// float型のURLクエリをセットする。
func QueryFloat[Float constraints.Float](key string, value Float) Option {
	return Query(key, strconv.FormatFloat(float64(value), 'f', -1, reflect.TypeOf(value).Bits()))
}
