package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Tag(mediaType MediaType) (TagResult, error) {
	requestUrl, _ := url.JoinPath(apiUrl, "/"+mediaType.Code+"/tag")
	var o TagResult
	err := c.get(requestUrl, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 tag 失败: %v", err)
	}
	return o, nil
}

type TagResult struct {
	ApiResult `mapstructure:",squash"`
	Name      string `json:"name" mapstructure:"name"`
	Tags      []Tag  `json:"tags" mapstructure:"tags"`
	//RelatedTags
	Filters          []Filter   `json:"filters" mapstructure:"filters"`
	ShowRatingFilter bool       `json:"showRatingFilter" mapstructure:"show_rating_filter"`
	Sorts            []Sort     `json:"sorts" mapstructure:"sorts"`
	Data             []RankItem `json:"data" mapstructure:"data"`
}

type Tag struct {
	Editable bool     `json:"editable" mapstructure:"editable"`
	Data     []string `json:"data" mapstructure:"data"`
	Type     string   `json:"type" mapstructure:"type"`
}
