package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Recommendations(mediaType MediaType, id string) ([]RecommendationsItem, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/recommendations")
	var o []RecommendationsItem
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 recommendations 失败: %v", err)
	}
	return o, nil
}

type RecommendationsItem struct {
	Rating     Rating  `json:"rating" mapstructure:"rating"`
	AlgJSON    string  `json:"alg_json" mapstructure:"alg_json"`
	SharingURL string  `json:"sharing_url" mapstructure:"sharing_url"`
	Title      string  `json:"title" mapstructure:"title"`
	URL        string  `json:"url" mapstructure:"url"`
	Pic        Pic     `json:"pic" mapstructure:"pic"`
	URI        string  `json:"uri" mapstructure:"uri"`
	Interest   *string `json:"interest" mapstructure:"interest"`
	Type       string  `json:"type" mapstructure:"type"`
	ID         string  `json:"id" mapstructure:"id"`
}
