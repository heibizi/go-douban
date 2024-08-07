package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

// Search 搜索 同时支持 movie tv 搜索，不需要指定 mediaType
func (c *ApiClient) Search(q string, start int, count int) (SearchResult, error) {
	requestUrl, _ := url.JoinPath(apiUrl, "/search/movie")
	var o SearchResult
	params := map[string]string{"q": q, "start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(requestUrl, params, &o)
	if err != nil {
		return o, fmt.Errorf("search 异常")
	}
	return o, nil
}

type SearchResult struct {
	ApiResult `mapstructure:",squash"`
	Banned    string       `json:"banned" mapstructure:"banned"`
	Items     []SearchItem `json:"items" mapstructure:"items"`
}
