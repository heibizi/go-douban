package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

type SubjectCollectionId string

const (
	MovieShowing SubjectCollectionId = "movie_showing"  // 正在热映
	MovieHotGaia SubjectCollectionId = "movie_hot_gaia" // 热门电影
	MovieSoon    SubjectCollectionId = "movie_soon"     // 即将上映
	MovieTop250  SubjectCollectionId = "movie_top250"   // Top250
	MovieSciFi   SubjectCollectionId = "movie_scifi"    // 高分经典科幻片
	MovieComeDyn SubjectCollectionId = "movie_comedy"   // 高分经典喜剧片
	MovieAction  SubjectCollectionId = "movie_action"   // 高分经典动作片
	MovieLove    SubjectCollectionId = "movie_love"     // 高分经典爱情片

	TvHot               SubjectCollectionId = "tv_hot"                 // 热门剧集
	TvDomestic          SubjectCollectionId = "tv_domestic"            // 国产剧
	TvAmerican          SubjectCollectionId = "tv_american"            // 美剧
	TvJapanese          SubjectCollectionId = "tv_japanese"            // 日剧
	TvKorean            SubjectCollectionId = "tv_korean"              // 韩剧
	TvAnimation         SubjectCollectionId = "tv_animation"           // 动漫
	TvVarietyShow       SubjectCollectionId = "tv_variety_show"        // 综艺
	TvChineseBestWeekly SubjectCollectionId = "tv_chinese_best_weekly" // 华语口碑周榜
	TvGlobalBestWeekly  SubjectCollectionId = "tv_global_best_weekly"  // 全球口碑周榜

	ShowHot      SubjectCollectionId = "show_hot"      // 执门综艺
	ShowDomestic SubjectCollectionId = "show_domestic" // 国内综艺
	ShowForeign  SubjectCollectionId = "show_foreign"  // 国外综艺
)

func (c *ApiClient) SubjectCollectionItems(id string, start int, count int) (SubjectCollectionItemsResult, error) {
	u, _ := url.JoinPath(apiUrl, "/subject_collection/", id, "/items")
	var o SubjectCollectionItemsResult
	params := map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 subject collection items 失败: %v", err)
	}
	return o, nil
}

func (c *ApiClient) SubjectCollectionItemsWith(id SubjectCollectionId, start int, count int) (SubjectCollectionItemsResult, error) {
	return c.SubjectCollectionItems(string(id), start, count)
}

type SubjectCollectionItemsResult struct {
	ApiResult              `mapstructure:",squash"`
	SubjectCollectionItems []SubjectCollectionItem `json:"subjectCollectionItems" mapstructure:"subject_collection_items"`
	SubjectCollection      SubjectCollectionResult `json:"subjectCollection" mapstructure:"subject_collection"`
}

type HonorInfo struct {
	Kind  string `json:"kind" mapstructure:"kind"`
	URI   string `json:"uri" mapstructure:"uri"`
	Rank  int    `json:"rank" mapstructure:"rank"`
	Title string `json:"title" mapstructure:"title"`
}

type SubjectCollectionItem struct {
	Comment           string      `json:"comment" mapstructure:"comment"`
	Rating            Rating      `json:"rating" mapstructure:"rating"`
	ControversyReason string      `json:"controversyReason" mapstructure:"controversy_reason"`
	Pic               Pic         `json:"pic" mapstructure:"pic"`
	Rank              int         `json:"rank" mapstructure:"rank"`
	URI               string      `json:"uri" mapstructure:"uri"`
	IsShow            bool        `json:"isShow" mapstructure:"is_show"`
	VendorIcons       []string    `json:"vendorIcons" mapstructure:"vendor_icons"`
	CardSubtitle      string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	ID                string      `json:"id" mapstructure:"id"`
	Title             string      `json:"title" mapstructure:"title"`
	HasLineWatch      bool        `json:"hasLineWatch" mapstructure:"has_linewatch"`
	IsReleased        bool        `json:"isReleased" mapstructure:"is_released"`
	Interest          interface{} `json:"interest" mapstructure:"interest"`
	ColorScheme       ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	Type              string      `json:"type" mapstructure:"type"`
	Description       string      `json:"description" mapstructure:"description"`
	Tags              []string    `json:"tags" mapstructure:"tags"`
	CoverURL          string      `json:"coverUrl" mapstructure:"cover_url"`
	Photos            []string    `json:"photos" mapstructure:"photos"`
	Actions           []string    `json:"actions" mapstructure:"actions"`
	SharingURL        string      `json:"sharingUrl" mapstructure:"sharing_url"`
	URL               string      `json:"url" mapstructure:"url"`
	HonorInfos        []HonorInfo `json:"honorInfos" mapstructure:"honor_infos"`
	GoodRatingStats   int         `json:"goodRatingStats" mapstructure:"good_rating_stats"`
	Subtype           string      `json:"subtype" mapstructure:"subtype"`
	NullRatingReason  string      `json:"nullRatingReason" mapstructure:"null_rating_reason"`
}
