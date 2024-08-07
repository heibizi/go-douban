package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Reviews(mediaType MediaType, id string) (ReviewsResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/reviews")
	var o ReviewsResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 reviews 失败: %v", err)
	}
	return o, nil
}

type ReviewsResult struct {
	ApiResult `mapstructure:",squash"`
	Reviews   []Review `json:"reviews" mapstructure:"reviews"`
}

type Review struct {
	Rating         Rating        `json:"rating" mapstructure:"rating"`
	UsefulCount    int           `json:"usefulCount" mapstructure:"useful_count"`
	SharingURL     string        `json:"sharingUrl" mapstructure:"sharing_url"`
	Title          string        `json:"title" mapstructure:"title"`
	URL            string        `json:"url" mapstructure:"url"`
	Abstract       string        `json:"abstract" mapstructure:"abstract"`
	URI            string        `json:"uri" mapstructure:"uri"`
	AdInfo         string        `json:"adInfo" mapstructure:"ad_info"`
	Topic          string        `json:"topic" mapstructure:"topic"`
	Photos         []ReviewPhoto `json:"photos" mapstructure:"photos"`
	ReactionsCount int           `json:"reactionsCount" mapstructure:"reactions_count"`
	CommentsCount  int           `json:"commentsCount" mapstructure:"comments_count"`
	User           User          `json:"user" mapstructure:"user"`
	CreateTime     string        `json:"createTime" mapstructure:"create_time"`
	ReSharesCount  int           `json:"reSharesCount" mapstructure:"reshares_count"`
	Type           string        `json:"type" mapstructure:"type"`
	ID             string        `json:"id" mapstructure:"id"`
	Subject        struct {
		Type  string `json:"type" mapstructure:"type"`
		Title string `json:"title" mapstructure:"title"`
	} `json:"subject" mapstructure:"subject"`
}

type ReviewPhoto struct {
	TagName     string `json:"tagName" mapstructure:"tag_name"`
	ID          string `json:"id" mapstructure:"id"`
	Image       Image  `json:"image" mapstructure:"image"`
	Description string `json:"description" mapstructure:"description"`
}
