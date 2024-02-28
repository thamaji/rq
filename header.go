package rq

import (
	"encoding/base64"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// string型のHTTPリクエストヘッダをセットする。
func Header[String ~string](key string, value String) Option {
	return OptionFunc(func(r *Request) {
		r.header.Set(key, string(value))
	})
}

// int型のHTTPリクエストヘッダをセットする。
func HeaderInt[Int ~int | ~int8 | ~int16 | ~int32 | ~int64](key string, value Int) Option {
	return Header(key, strconv.FormatInt(int64(value), 10))
}

// uint型のHTTPリクエストヘッダをセットする。
func HeaderUint[Uint ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr](key string, value Uint) Option {
	return Header(key, strconv.FormatUint(uint64(value), 10))
}

// bool型のHTTPリクエストヘッダをセットする。
func HeaderBool[Bool ~bool](key string, value Bool) Option {
	return Header(key, strconv.FormatBool(bool(value)))
}

// float型のHTTPリクエストヘッダをセットする。
func HeaderFloat[Float ~float32 | ~float64](key string, value Float) Option {
	return Header(key, strconv.FormatFloat(float64(value), 'f', -1, reflect.TypeOf(value).Bits()))
}

type queryValue float64

func (q queryValue) String() string {
	return strconv.FormatFloat(float64(q), 'f', 2, 64)
}

// Acceptヘッダをセットする。
func Accept(mediaType ...string) AcceptOption {
	return AcceptOption{strings.Join(mediaType, ", ")}
}

type AcceptOption struct {
	v string
}

func (option AcceptOption) Apply(r *Request) {
	r.header.Set("Accept", option.v)
}

func (option AcceptOption) MediaType(mediaType string, q float64) AcceptOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += mediaType
	if q != 0 {
		v += ";" + queryValue(q).String()
	}
	return AcceptOption{v}
}

// Accept-Charsetヘッダをセットする。
func AcceptCharset(charset ...string) AcceptCharsetOption {
	return AcceptCharsetOption{strings.Join(charset, ", ")}
}

type AcceptCharsetOption struct {
	v string
}

func (option AcceptCharsetOption) Apply(r *Request) {
	r.header.Set("Accept-Charset", option.v)
}

func (option AcceptCharsetOption) Charset(charset string, q float64) AcceptCharsetOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += charset
	if q != 0 {
		v += ";" + queryValue(q).String()
	}
	return AcceptCharsetOption{v}
}

// Accept-Encodingヘッダをセットする。
func AcceptEncoding(encoding ...string) AcceptEncodingOption {
	return AcceptEncodingOption{strings.Join(encoding, ", ")}
}

type AcceptEncodingOption struct {
	v string
}

func (option AcceptEncodingOption) Apply(r *Request) {
	r.header.Set("Accept-Encoding", option.v)
}

func (option AcceptEncodingOption) Encoding(encoding string, q float64) AcceptEncodingOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += encoding
	if q != 0 {
		v += ";" + queryValue(q).String()
	}
	return AcceptEncodingOption{v}
}

// Accept-Languageヘッダをセットする。
func AcceptLanguage(language ...string) AcceptLanguageOption {
	return AcceptLanguageOption{strings.Join(language, ", ")}
}

type AcceptLanguageOption struct {
	v string
}

func (option AcceptLanguageOption) Apply(r *Request) {
	r.header.Set("Accept-Language", option.v)
}

func (option AcceptLanguageOption) Language(language string, q float64) AcceptLanguageOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += language
	if q != 0 {
		v += ";" + queryValue(q).String()
	}
	return AcceptLanguageOption{v}
}

// Allowヘッダをセットする。
func Allow(method ...string) Option {
	return Header("Allow", strings.Join(method, ", "))
}

// Authorizationヘッダをセットする。
func Authorization(type_ string, credentials string) Option {
	return Header("Authorization", type_+" "+credentials)
}

// Basic認証のAuthorizationヘッダをセットする。
func AuthorizationBasic(username string, password string) Option {
	return Authorization("Basic", base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
}

// Bearer認証のAuthorizationヘッダをセットする。
func AuthorizationBearer(token string) Option {
	return Authorization("Bearer", token)
}

// Cache-Controlヘッダをセットする。
func CacheControl(control ...string) CacheControlOption {
	return CacheControlOption{strings.Join(control, ", ")}
}

type CacheControlOption struct {
	v string
}

func (option CacheControlOption) Apply(r *Request) {
	r.header.Set("Cache-Control", option.v)
}

func (option CacheControlOption) NoCache() CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "no-cache"
	return CacheControlOption{v}
}

func (option CacheControlOption) NoStore() CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "no-store"
	return CacheControlOption{v}
}

func (option CacheControlOption) MaxAge(maxAge time.Duration) CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "max-age=" + strconv.Itoa(int(maxAge.Seconds()))
	return CacheControlOption{v}
}

func (option CacheControlOption) MaxStale(maxStale time.Duration) CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "max-stale=" + strconv.Itoa(int(maxStale.Seconds()))
	return CacheControlOption{v}
}

func (option CacheControlOption) MaxFresh(maxFresh time.Duration) CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "max-fresh=" + strconv.Itoa(int(maxFresh.Seconds()))
	return CacheControlOption{v}
}

func (option CacheControlOption) NoTransform() CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "no-transform"
	return CacheControlOption{v}
}

func (option CacheControlOption) OnlyIfCached() CacheControlOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	v += "only-if-cached"
	return CacheControlOption{v}
}

// Content-Encodingヘッダをセットする。
func ContentEncoding(encoding ...string) Option {
	return Header("Content-Encoding", strings.Join(encoding, ", "))
}

// Content-Languageヘッダをセットする。
func ContentLanguage(language ...string) Option {
	return Header("Content-Language", strings.Join(language, ", "))
}

// Content-Lengthヘッダをセットする。
func ContentLength(bytes int) Option {
	return HeaderInt("Content-Length", bytes)
}

// Content-Locationヘッダをセットする。
func ContentLocation(location string) Option {
	return Header("Content-Location", location)
}

// Content-Typeヘッダをセットする。
func ContentType(mediaType string) ContentTypeOption {
	return ContentTypeOption{mediaType}
}

type ContentTypeOption struct {
	v string
}

func (option ContentTypeOption) Apply(r *Request) {
	r.header.Set("Content-Type", option.v)
}

func (option ContentTypeOption) Charset(charset string) ContentTypeOption {
	v := option.v
	v += "; charset=" + charset
	return ContentTypeOption{v}
}

func (option ContentTypeOption) Boundary(boundary string) ContentTypeOption {
	v := option.v
	v += "; boundary=" + boundary
	return ContentTypeOption{v}
}

// Cookieヘッダをセットする。
func Cookie(cookie ...http.Cookie) Option {
	v := make([]string, 0, len(cookie))
	for i := range cookie {
		v = append(v, cookie[i].String())
	}
	return Header("Cookie", strings.Join(v, "; "))
}

// Expectヘッダをセットする。
func Expect(expect string) Option {
	return Header("Expect", expect)
}

// Forwardedヘッダをセットする。
func Forwarded(by string, for_ string, host string, proto string) ForwardedOption {
	return ForwardedOption{}.And(by, for_, host, proto)
}

type ForwardedOption struct {
	v string
}

func (option ForwardedOption) Apply(r *Request) {
	r.header.Set("Forwarded", option.v)
}

func (option ForwardedOption) And(by string, for_ string, host string, proto string) ForwardedOption {
	v := []string{}
	if by != "" {
		v = append(v, "by="+by)
	}
	if for_ != "" {
		v = append(v, "for="+for_)
	}
	if host != "" {
		v = append(v, "host="+host)
	}
	if proto != "" {
		v = append(v, "proto="+proto)
	}
	if option.v == "" {
		return ForwardedOption{strings.Join(v, ";")}
	}
	return ForwardedOption{option.v + " ," + strings.Join(v, ";")}
}

// Fromヘッダをセットする。
func From(email string) Option {
	return Header("From", email)
}

// Hostヘッダをセットする。
func Host(host string) HostOption {
	return HostOption{host}
}

type HostOption struct {
	v string
}

func (option HostOption) Apply(r *Request) {
	r.header.Set("Host", option.v)
}

func (option HostOption) Port(port int) HostOption {
	v := option.v
	v += ":" + strconv.Itoa(port)
	return HostOption{v}
}

// If-Matchヘッダをセットする。
func IfMatch(etagValue ...string) Option {
	values := make([]string, len(etagValue))
	for i := range etagValue {
		if etagValue[i] == "*" {
			values[i] = "*"
		} else {
			values[i] = "\"" + etagValue[i] + "\""
		}
	}
	return Header("If-Match", strings.Join(values, ", "))
}

// If-Modified-Sinceヘッダをセットする。
func IfModifiedSince(timestamp time.Time) Option {
	return Header("If-Modified-Since", timestamp.Format(time.RFC1123))
}

// If-None-Matchヘッダをセットする。
func IfNoneMatch(etagValue ...string) Option {
	values := make([]string, len(etagValue))
	for i := range etagValue {
		if etagValue[i] == "*" {
			values[i] = "*"
		} else {
			values[i] = "\"" + etagValue[i] + "\""
		}
	}
	return Header("If-None-Match", strings.Join(values, ", "))
}

// If-Rangeヘッダをセットする。
func IfRange(timestamp time.Time, etagValue string) Option {
	return OptionFunc(func(r *Request) {
		if etagValue != "*" {
			etagValue = "\"" + etagValue + "\""
		}
		r.header.Set("If-Range", timestamp.Format(time.RFC1123))
		r.header.Add("If-Range", etagValue)
	})
}

// If-Unmodified-Sinceヘッダをセットする。
func IfUnmodifiedSince(timestamp time.Time) Option {
	return Header("If-Unmodified-Since", timestamp.Format(time.RFC1123))
}

// Originヘッダをセットする。
func Origin(origin string) Option {
	return Header("Origin", origin)
}

// Proxy-Authorizationヘッダをセットする。
func ProxyAuthorization(type_ string, credentials string) Option {
	return Header("Proxy-Authorization", type_+" "+credentials)
}

// Basic認証のProxy-Authorizationヘッダをセットする。
func ProxyAuthorizationBasic(username string, password string) Option {
	return ProxyAuthorization("Basic", base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
}

// Bearer認証のProxy-Authorizationヘッダをセットする。
func ProxyAuthorizationBearer(token string) Option {
	return ProxyAuthorization("Bearer", token)
}

// Rangeヘッダをセットする。
func Range(start int, end int) RangeOption {
	return RangeOption{v: "bytes="}.And(start, end)
}

type RangeOption struct {
	v string
}

func (option RangeOption) Apply(r *Request) {
	r.header.Set("Host", option.v)
}

func (option RangeOption) And(start int, end int) RangeOption {
	v := option.v
	if v != "" {
		v += ", "
	}
	if start != 0 {
		v += strconv.Itoa(start)
	}
	v += "-"
	if end != 0 {
		v += strconv.Itoa(end)
	}
	return RangeOption{v}
}

// Refererヘッダをセットする。
func Referer(referer string) Option {
	return Header("Referer", referer)
}

// User-Agentヘッダをセットする。
func UserAgent(userAgent string) Option {
	return Header("User-Agent", userAgent)
}

// Viaヘッダをセットする。
func Via(via string) Option {
	return Header("Via", via)
}

// X-Forwarded-Forヘッダをセットする。
func XForwardedFor(client string, proxy ...string) Option {
	return Header("X-Forwarded-For", strings.Join(append([]string{client}, proxy...), ", "))
}

// X-Forwarded-Hostヘッダをセットする。
func XForwardedHost(host string) Option {
	return Header("X-Forwarded-Host", host)
}

// X-Forwarded-Protoヘッダをセットする。
func XForwardedProto(protocol string) Option {
	return Header("X-Forwarded-Proto", protocol)
}
