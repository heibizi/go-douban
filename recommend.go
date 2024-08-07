package douban

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (c *ApiClient) Recommend(mediaType MediaType, tags []string, sort string, start int, count int) (RecommendResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/recommend")
	var o RecommendResult
	params := map[string]string{"sort": sort, "start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	if len(tags) > 0 {
		params["tags"] = strings.Join(tags, ",")
	}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 recommend 失败: %v", err)
	}
	var chartItems []RecommendChartItem
	for _, item := range o.ChartItems {
		if item.Type == "chart" {
			chartItems = append(chartItems, item)
		}
	}
	var items []RecommendItem
	for _, item := range o.Items {
		if item.Type != "chart" {
			items = append(items, item)
		}
	}
	o.ChartItems = chartItems
	o.Items = items
	return o, nil
}

type (
	RecommendResult struct {
		ApiResult                `mapstructure:",squash"`
		ShowRatingFilter         bool                     `json:"showRatingFilter" mapstructure:"show_rating_filter"`
		MovieRecommendCategories []MovieRecommendCategory `json:"movieRecommendCategories" mapstructure:"recommend_categories"`
		TvRecommendCategories    []TvRecommendCategory    `json:"tvRecommendCategories" mapstructure:"recommend_categories"`
		ChartItems               []RecommendChartItem     `json:"chartItems" mapstructure:"items"`
		Items                    []RecommendItem          `json:"items" mapstructure:"items"`
		BottomRecommendTags      []string                 `json:"bottomRecommendTags" mapstructure:"bottom_recommend_tags"`
		PlayableFilters          []PlayableFilter         `json:"playableFilters" mapstructure:"playable_filters"`
		Filters                  []Filter                 `json:"filters" mapstructure:"filters"`
		//QuickMark
		RecommendTags []string `json:"recommendTags" mapstructure:"recommend_tags"`
		//ManualTags
		SetUserVendor SetUserVendor `json:"setUserVendor" mapstructure:"set_user_vendor"`
		Sorts         []Sort        `json:"sorts" mapstructure:"sorts"`
	}
)

type SetUserVendor struct {
	URI   string `json:"uri" mapstructure:"uri"`
	Key   string `json:"key" mapstructure:"key"`
	Title string `json:"title" mapstructure:"title"`
}

type MovieRecommendCategory struct {
	IsControl bool   `json:"isControl" mapstructure:"is_control"`
	Type      string `json:"type" mapstructure:"type"`
	Data      []struct {
		Default bool   `json:"default" mapstructure:"default"`
		Text    string `json:"text" mapstructure:"text"`
	} `json:"data" mapstructure:"data"`
}

type TvRecommendCategory struct {
	IsControl bool   `json:"isControl" mapstructure:"is_control"`
	Type      string `json:"type" mapstructure:"type"`
	Data      []struct {
		Default   bool     `json:"default" mapstructure:"default"`
		IsControl bool     `json:"isControl" mapstructure:"is_control"`
		Text      string   `json:"text" mapstructure:"text"`
		Tags      []string `json:"tags" mapstructure:"tags"`
	} `json:"data" mapstructure:"data"`
	TagGroups string `json:"tagGroups" mapstructure:"tag_groups"`
}

type RecommendChartItem struct {
	ItemsCount     int         `json:"itemsCount" mapstructure:"items_count"`
	Subtitle       string      `json:"subtitle" mapstructure:"subtitle"`
	Title          string      `json:"title" mapstructure:"title"`
	IsFollow       bool        `json:"isFollow" mapstructure:"is_follow"`
	UseImageCover  bool        `json:"useImageCover" mapstructure:"use_image_cover"`
	URI            string      `json:"uri" mapstructure:"uri"`
	CoverURL       string      `json:"coverURL" mapstructure:"cover_url"`
	ItemType       string      `json:"itemType" mapstructure:"item_type"`
	FollowersCount int         `json:"followersCount" mapstructure:"followers_count"`
	BGImage        string      `json:"bgImage" mapstructure:"bg_image"`
	TileCover      bool        `json:"tileCover" mapstructure:"tile_cover"`
	AlgJSON        string      `json:"algJson" mapstructure:"alg_json"`
	ColorScheme    ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	Type           string      `json:"type" mapstructure:"type"`
	ID             string      `json:"id" mapstructure:"id"`
	Card           string      `json:"card" mapstructure:"card"`
}

type Comment struct {
	Comment string `json:"comment" mapstructure:"comment"`
	ID      string `json:"id" mapstructure:"id"`
	User    Owner  `json:"user" mapstructure:"user"`
}

type Owner struct {
	Kind   string `json:"kind" mapstructure:"kind"`
	Name   string `json:"name" mapstructure:"name"`
	URL    string `json:"url" mapstructure:"url"`
	URI    string `json:"uri" mapstructure:"uri"`
	Avatar string `json:"avatar" mapstructure:"avatar"`
	IsClub bool   `json:"isClub" mapstructure:"is_club"`
	Type   string `json:"type" mapstructure:"type"`
	ID     string `json:"id" mapstructure:"id"`
	UID    string `json:"uid" mapstructure:"uid"`
}

type RecommendItem struct {
	Comment      Comment     `json:"comment" mapstructure:"comment"`
	Rating       Rating      `json:"rating" mapstructure:"rating"`
	VendorCount  int         `json:"vendorCount" mapstructure:"vendor_count"`
	PlayableDate *string     `json:"playableDate" mapstructure:"playable_date"`
	Pic          Picture     `json:"pic" mapstructure:"pic"`
	HonorInfos   []HonorInfo `json:"honorInfos" mapstructure:"honor_infos"`
	VendorIcons  []string    `json:"vendorIcons" mapstructure:"vendor_icons"`
	Year         string      `json:"year" mapstructure:"year"`
	CardSubtitle string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	ID           string      `json:"id" mapstructure:"id"`
	Title        string      `json:"title" mapstructure:"title"`
	Tags         []struct {
		Name string `json:"name" mapstructure:"name"`
		URI  string `json:"uri" mapstructure:"uri"`
	} `json:"tags" mapstructure:"tags"`
	Interest         *string  `json:"interest" mapstructure:"interest"`
	Type             string   `json:"type" mapstructure:"type"`
	AlgJSON          string   `json:"algJson" mapstructure:"alg_json"`
	HasLineWatch     bool     `json:"hasLineWatch" mapstructure:"has_linewatch"`
	Photos           []string `json:"photos" mapstructure:"photos"`
	Card             string   `json:"card" mapstructure:"card"`
	PlayableDateInfo string   `json:"playableDateInfo" mapstructure:"playable_date_info"`
	FollowingRating  *string  `json:"followingRating" mapstructure:"following_rating"`
	URI              string   `json:"uri" mapstructure:"uri"`
	EpisodesInfo     string   `json:"episodesInfo" mapstructure:"episodes_info"`
	ItemType         string   `json:"itemType" mapstructure:"item_type"`
}

type PlayableFilter struct {
	URI   string `json:"uri" mapstructure:"uri"`
	Key   string `json:"key" mapstructure:"key"`
	Title string `json:"title" mapstructure:"title"`
}

type Filter struct {
	Text    string `json:"text" mapstructure:"text"`
	Checked bool   `json:"checked" mapstructure:"checked"`
	Name    string `json:"name" mapstructure:"name"`
	Desc    string `json:"desc" mapstructure:"desc"`
}

type Sort struct {
	Text    string `json:"text" mapstructure:"text"`
	Checked bool   `json:"checked" mapstructure:"checked"`
	Name    string `json:"name" mapstructure:"name"`
}
