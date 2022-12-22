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

