package douban

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"io"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const apiKey = "0dad551ec0f84ed02907ff5c42e8ec70"
const apiUrl = "https://frodo.douban.com/api/v2"
const secretKey = "bf7dddc7c9cfe6f7"

var userAgents = []string{
	"api-client/1 com.douban.frodo/7.22.0.beta9(231) Android/23 product/Mate 40 vendor/HUAWEI model/Mate 40 brand/HUAWEI  rom/android  network/wifi  platform/AndroidPad",
	"api-client/1 com.douban.frodo/7.18.0(230) Android/22 product/MI 9 vendor/Xiaomi model/MI 9 brand/Android  rom/miui6  network/wifi  platform/mobile nd/1",
	"api-client/1 com.douban.frodo/7.1.0(205) Android/29 product/perseus vendor/Xiaomi model/Mi MIX 3  rom/miui6  network/wifi  platform/mobile nd/1",
	"api-client/1 com.douban.frodo/7.3.0(207) Android/22 product/MI 9 vendor/Xiaomi model/MI 9 brand/Android  rom/miui6  network/wifi platform/mobile nd/1",
}

type ApiClient struct {
}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (c *ApiClient) header() map[string]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return map[string]string{
		"User-Agent": userAgents[r.Intn(len(userAgents))],
	}
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
	resp, err := HttpRequest(url, method, c.header(), p, nil)
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
