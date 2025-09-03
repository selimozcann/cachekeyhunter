package scan

import (
	urlpkg "net/url"

	"github.com/selimozcann/cachekeyhunter/ck/internal/httpx"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func GetBaseline(target string) (types.Signals, error) {
	return httpx.DoRequest(target, nil)
}

func DoVariant(target string, v types.Variant) (types.Signals, error) {
	u, err := urlpkg.Parse(target)
	if err != nil {
		return types.Signals{}, err
	}
	if len(v.Query) > 0 {
		q := u.Query()
		for k, val := range v.Query {
			q.Set(k, val)
		}
		u.RawQuery = q.Encode()
	}
	return httpx.DoRequest(u.String(), v.Headers)
}
