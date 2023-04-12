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

type Params map[string]string

type baseClient struct {
	baseUrl   string
	client    *http.Client
	apiKey    string
	apiSecret []byte
}

func newBaseClient(baseUrl, apiKey, apiSecret string) (c *baseClient) {
	return &baseClient{
		baseUrl:   baseUrl,
		client:    &http.Client{},
		apiKey:    apiKey,
		apiSecret: []byte(apiSecret),
	}
}

func (c *baseClient) Get(path string, signed bool, params Params, resultStructPtr interface{}) (err error) {
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
			if errStruct.Code == -1022 {
				delete(params, "signature")
				println(errStruct.Message)
				return c.Get(path, signed, params, resultStructPtr)
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
