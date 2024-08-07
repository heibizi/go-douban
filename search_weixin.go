package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) SearchWx(q string, start int, count int) (WxResult, error) {
	u, _ := url.JoinPath(apiUrl, "/search/weixin")
	var o WxResult
	params := map[string]string{"q": q, "start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("search weixin 异常")
	}
	var items []SearchItem
	var douListItems []SearchDouListItem
	for _, item := range o.Items {
		if item.TargetType == "move" || item.TargetType == "tv" {
			items = append(items, item)
		}
	}
	for _, item := range o.DouListItems {
		if item.TargetType == "doulist_cards" {
			douListItems = append(douListItems, item)
		}
	}
	o.Items = items
	o.DouListItems = douListItems
	return o, nil
}

type WxResult struct {
	ApiResult    `mapstructure:",squash"`
	Items        []SearchItem        `json:"items" mapstructure:"items"`
	DouListItems []SearchDouListItem `json:"douListItems" mapstructure:"items"`
}

type Target struct {
	Rating            Rating `json:"rating" mapstructure:"rating"`
	ControversyReason string `json:"controversyReason" mapstructure:"controversy_reason"`
	Title             string `json:"title" mapstructure:"title"`
	Abstract          string `json:"abstract" mapstructure:"abstract"`
	HasLineWatch      bool   `json:"hasLineWatch" mapstructure:"has_linewatch"`
	URI               string `json:"uri" mapstructure:"uri"`
	CoverURL          string `json:"coverUrl" mapstructure:"cover_url"`
	Year              string `json:"year" mapstructure:"year"`
	CardSubtitle      string `json:"cardSubtitle" mapstructure:"card_subtitle"`
	ID                string `json:"id" mapstructure:"id"`
	NullRatingReason  string `json:"nullRatingReason" mapstructure:"null_rating_reason"`
}

type SearchItem struct {
	Layout     string `json:"layout" mapstructure:"layout"`
	TypeName   string `json:"typeName" mapstructure:"type_name"`
	TargetID   string `json:"targetId" mapstructure:"target_id"`
	Target     Target `json:"target" mapstructure:"target"`
	TargetType string `json:"targetType" mapstructure:"target_type"`
}

type SearchDouListItem struct {
	Title        string `json:"title" mapstructure:"title"`
	IsBadgeChart bool   `json:"isBadgeChart" mapstructure:"is_badge_chart"`
	URI          string `json:"uri" mapstructure:"uri"`
	CoverURL     string `json:"coverUrl" mapstructure:"cover_url"`
	TargetType   string `json:"targetType" mapstructure:"target_type"`
	CardSubtitle string `json:"cardSubtitle" mapstructure:"card_subtitle"`
	IsOfficial   bool   `json:"isOfficial" mapstructure:"is_official"`
	ID           string `json:"id" mapstructure:"id"`
	ImageLabel   string `json:"imageLabel" mapstructure:"image_label"`
}
