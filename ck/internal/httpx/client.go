package httpx

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	Constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

var httpClient = &http.Client{
	Timeout: Constants.DefaultTimeoutSeconds * time.Second,
}

func DoRequest(url string, headers map[string]string) (types.Signals, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set(Constants.UserAgentHeader, Constants.UserAgentValue)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return types.Signals{}, err
	}
	defer resp.Body.Close()

	lim := io.LimitReader(resp.Body, 512<<10)
	h := sha256.New()
	io.Copy(h, lim)

	s := types.Signals{
		StatusCode: resp.StatusCode, Headers: map[string]string{}, BodyHash: hex.EncodeToString(h.Sum(nil)),
	}

	for _, k := range []string{Constants.HeaderXCache, Constants.HeaderCFCacheStatus, Constants.HeaderAge, Constants.HeaderVary} {
		v := resp.Header.Get(k)
		s.Headers[k] = v
		if k == Constants.HeaderAge {
			n, _ := strconv.Atoi(v)
			s.Age = n
		}
	}
	if s.Cache == "" {
		s.Cache = resp.Header.Get(Constants.HeaderXCache)
		if s.Cache == "" {
			s.Cache = resp.Header.Get(Constants.HeaderCFCacheStatus)
		}
	}
	return s, nil

}
