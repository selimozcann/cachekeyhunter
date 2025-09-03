package scan

import (
	"github.com/selimozcann/cachekeyhunter/ck/internal/httpx"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func GetBaseLİne(url string) (types.Signals, error) {
	return httpx.DoRequest(url, nil)
}

func DoVariant(url string, v types.Variant) (types.Signals, error) {
	return httpx.DoRequest(url, v.Headers)
}
