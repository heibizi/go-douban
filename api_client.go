package douban

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"io"
	"net/url"
	"strings"
	"time"
)

const apiKey = "054022eaeae0b00e0fc068c0c0a2102a"
const apiUrl = "https://frodo.douban.com/api/v2"
const secretKey = "bf7dddc7c9cfe6f7"

var header = map[string]string{
	"User-Agent": "MicroMessenger/",
	"Referer":    "https://servicewechat.com/wx2f9b06c1de1ccfca/91/page-frame.html",
}

type ApiClient struct {
}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (c *ApiClient) sign(urlStr string, ts string, method string) string {
	parsedURL, _ := url.Parse(urlStr)
	urlPath := parsedURL.Path
	rawSign := strings.Join([]string{strings.ToUpper(method), url.QueryEscape(urlPath), ts}, "&")
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(rawSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

func (c *ApiClient) request(url string, method string, params map[string]string) ([]byte, error) {
	p := make(map[string]string)
	p["apiKey"] = apiKey
	p["os_rom"] = "android"
	ts := time.Now().Format("20060102")
	p["_sig"] = c.sign(url, ts, method)
	p["_ts"] = ts
	if params != nil {
		for k, v := range params {
			p[k] = v
		}
	}
	resp, err := HttpRequest(url, method, header, p, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *ApiClient) result(data []byte, o interface{}) error {
	result := gjson.Parse(string(data))
	err := mapstructure.WeakDecode(result.Value(), &o)
	if err != nil {
		return err
	}
	return nil
}

func (c *ApiClient) get(url string, params map[string]string, o interface{}) error {
	data, err := c.request(url, "GET", params)
	if err != nil {
		return err
	}
	return c.result(data, o)
}
