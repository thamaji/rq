package rq

type Option interface {
	Apply(*Request)
}

type OptionFunc func(*Request)

func (f OptionFunc) Apply(r *Request) {
	f(r)
}
