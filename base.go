package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var maxAttemptsCount = 5

type Params map[string]string

type baseClient struct {
	baseUrl   string
	client    *http.Client
	apiKey    string
	apiSecret []byte
	proxy     *Proxy
}

func newBaseClient(baseUrl, apiKey, apiSecret string) (c *baseClient) {
	if defaultProxy == nil {
		defaultProxy = NewProxy()
	}
	return &baseClient{
		baseUrl:   baseUrl,
		client:    &http.Client{},
		apiKey:    apiKey,
		apiSecret: []byte(apiSecret),
		proxy:     defaultProxy,
	}
}

func (c *baseClient) Get(path string, signed bool, params Params, resultStructPtr interface{}, weight int) (err error) {
	c.proxy.waitForWeight(weight)
	attempt := 0
	return c.get(path, signed, params, resultStructPtr, attempt)
}

func (c *baseClient) get(path string, signed bool, params Params, resultStructPtr interface{}, attempt int) (err error) {
	query := c.getQuery(params)
	if signed {
		query = c.signQuery(query)
	}
	req, err := http.NewRequest("GET", c.getUrl(path, query), nil)
	if err != nil {
		err = errors.Wrap(err, "request error")
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-MBX-APIKEY", c.apiKey)
	req.Header.Add("User-Agent", "oracle")
	resp, err := c.client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "response error")
		return
	}
	if resp.StatusCode != 200 {
		var errStruct binanceError
		if resultStructPtr != nil {
			if err = json.NewDecoder(resp.Body).Decode(&errStruct); err != nil {
				return errors.Wrap(err, "unmarshalling json error")
			}
			switch errStruct.Code {
			case -1021:
				delete(params, "signature")
				println(errStruct.Message)
				if attempt >= maxAttemptsCount {
					return errors.Wrap(errors.New(fmt.Sprintf("Code: %d Message: %s", errStruct.Code, errStruct.Message)), resp.Status)
				}
				attempt++
				return c.get(path, signed, params, resultStructPtr, attempt)
			case -1022:
				return ErrorBadAPISecret
			case -2014, -2015:
				return ErrorBadAPIKey
			}
			return errors.Wrap(errors.New(fmt.Sprintf("Code: %d Message: %s", errStruct.Code, errStruct.Message)), resp.Status)
		}
	}
	if resultStructPtr != nil {
		if err = json.NewDecoder(resp.Body).Decode(resultStructPtr); err != nil {
			return errors.Wrap(err, "unmarshalling json error")
		}
	}
	return nil
}

func (c *baseClient) getQuery(p Params) (q string) {
	if p != nil && len(p) != 0 {
		buf := make([]string, 0, len(p))
		for k, v := range p {
			buf = append(buf, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
		q = strings.Join(buf, "&")
	}
	return
}

func (c *baseClient) signQuery(q string) string {
	// timestamp and signature must be in the end: ...&timestamp=xxxxxx&signature=xxxxxx
	if q != "" {
		q += "&"
	}
	q += "timestamp=" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	q = q + "&signature=" + c.getSignature(q)
	return q
}

func (c *baseClient) getUrl(path string, query string) (res string) {
	res = c.baseUrl + path
	if query != "" {
		res += "?" + query
	}
	return res
}

func (c *baseClient) getSignature(query string) string {
	sig := hmac.New(sha256.New, c.apiSecret)
	sig.Write([]byte(query))
	return hex.EncodeToString(sig.Sum(nil))
}
