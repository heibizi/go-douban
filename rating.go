package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Rating(mediaType MediaType, id string) (RatingResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/rating")
	var o RatingResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 rating 失败: %v", err)
	}
	return o, nil
}

type WishFriends struct {
	Total int           `json:"total" mapstructure:"total"`
	Users []interface{} `json:"users" mapstructure:"users"`
}

type TypeRank struct {
	Type string  `json:"type" mapstructure:"type"`
	Rank float64 `json:"rank" mapstructure:"rank"`
}

type RatingResult struct {
	Stats       []float64   `json:"stats" mapstructure:"stats"`
	DoingCount  int         `json:"doingCount" mapstructure:"doing_count"`
	WishCount   int         `json:"wishCount" mapstructure:"wish_count"`
	WishFriends WishFriends `json:"wishFriends" mapstructure:"wish_friends"`
	TypeRanks   []TypeRank  `json:"typeRanks" mapstructure:"type_ranks"`
	Following   interface{} `json:"following" mapstructure:"following"`
	DoneCount   int         `json:"doneCount" mapstructure:"done_count"`
}
