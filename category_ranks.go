package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) CategoryRanks(mediaType MediaType, start int, count int) (CategoryRanksResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/category_ranks")
	var o CategoryRanksResult
	params := map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 category_ranks 失败: %v", err)
	}
	return o, nil
}

type CategoryRanksResult struct {
	ApiResult           `mapstructure:",squash"`
	SelectedCollections []SelectedCollection `json:"selectedCollections" mapstructure:"selected_collections"`
}
