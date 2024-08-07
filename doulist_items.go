package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) DouListItems(id string, start int, count int) (DouListItemsResult, error) {
	u, _ := url.JoinPath(apiUrl, "/doulist/", id, "/items")
	var o DouListItemsResult
	params := map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 doulist items 失败: %v", err)
	}
	return o, nil
}

type DouListItemsResult struct {
	ApiResult `mapstructure:",squash"`
	Items     []DouListItem `json:"items" mapstructure:"items"`
}

type DouListItem struct {
	Comment    string `json:"comment" mapstructure:"comment"`
	Rating     Rating `json:"rating" mapstructure:"rating"`
	Subtitle   string `json:"subtitle" mapstructure:"subtitle"`
	Title      string `json:"title" mapstructure:"title"`
	URL        string `json:"url" mapstructure:"url"`
	TargetID   string `json:"targetId" mapstructure:"target_id"`
	URI        string `json:"uri" mapstructure:"uri"`
	CoverURL   string `json:"coverURL" mapstructure:"cover_url"`
	CreateTime string `json:"createTime" mapstructure:"create_time"`
	Type       string `json:"type" mapstructure:"type"`
	ID         string `json:"id" mapstructure:"id"`
}
