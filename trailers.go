package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Trailers(mediaType MediaType, id string) (TrailersResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/trailers")
	var o TrailersResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 trailers 失败: %v", err)
	}
	return o, nil
}

type TrailersResult struct {
	Trailers []TrailerItem `json:"trailers" mapstructure:"trailers"`
}

type TrailerItem struct {
	ReactionType     int    `json:"reactionType" mapstructure:"reaction_type"`
	SharingURL       string `json:"sharingUrl" mapstructure:"sharing_url"`
	VideoURL         string `json:"videoUrl" mapstructure:"video_url"`
	Title            string `json:"title" mapstructure:"title"`
	TypeName         string `json:"typeName" mapstructure:"type_name"`
	URI              string `json:"uri" mapstructure:"uri"`
	CoverURL         string `json:"coverUrl" mapstructure:"cover_url"`
	TermNum          int    `json:"termNum" mapstructure:"term_num"`
	NComments        int    `json:"nComments" mapstructure:"n_comments"`
	ReactionsCount   int    `json:"reactionsCount" mapstructure:"reactions_count"`
	CreateTime       string `json:"createTime" mapstructure:"create_time"`
	CollectionsCount int    `json:"collectionsCount" mapstructure:"collections_count"`
	FileSize         int    `json:"fileSize" mapstructure:"file_size"`
	ReSharesCount    int    `json:"reSharesCount" mapstructure:"reshares_count"`
	Runtime          string `json:"runtime" mapstructure:"runtime"`
	Type             string `json:"type" mapstructure:"type"`
	ID               string `json:"id" mapstructure:"id"`
	IsCollected      bool   `json:"isCollected" mapstructure:"is_collected"`
	Desc             string `json:"desc" mapstructure:"desc"`
}
