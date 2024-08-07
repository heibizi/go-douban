package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) YearRanks(mediaType MediaType, year int) (YearRanks, error) {
	var o YearRanks
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/year_ranks")
	params := map[string]string{"year": strconv.Itoa(year)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 year_ranks 失败: %v", err)
	}
	return o, nil
}

type YearRanks struct {
	Groups []struct {
		Title               string               `json:"title" mapstructure:"title"`
		Uri                 string               `json:"uri" mapstructure:"uri"`
		SelectedCollections []SelectedCollection `json:"selected_collections" mapstructure:"selected_collections"`
	} `json:"groups"`
}
