package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Photos(mediaType MediaType, id string) (PhotosResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/photos")
	var o PhotosResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 photos 失败: %v", err)
	}
	return o, nil
}

type PhotosResult struct {
	ApiResult `mapstructure:",squash"`
	C         int     `json:"c" mapstructure:"c"`
	F         int     `json:"f" mapstructure:"f"`
	O         int     `json:"o" mapstructure:"o"`
	N         int     `json:"n" mapstructure:"n"`
	W         int     `json:"w" mapstructure:"w"`
	Photos    []Photo `json:"photos" mapstructure:"photos"`
	//AdList
}

type ImageDetail struct {
	URL          string      `json:"url" mapstructure:"url"`
	Width        int         `json:"width" mapstructure:"width"`
	Height       int         `json:"height" mapstructure:"height"`
	Size         int         `json:"size" mapstructure:"size"`
	PrimaryColor string      `json:"primaryColor" mapstructure:"primary_color"`
	IsAnimated   bool        `json:"isAnimated" mapstructure:"is_animated"`
	Raw          interface{} `json:"raw" mapstructure:"raw"`
}

type Photo struct {
	ReadCount        int         `json:"readCount" mapstructure:"read_count"`
	Image            Image       `json:"image" mapstructure:"image"`
	CreateTime       string      `json:"createTime" mapstructure:"create_time"`
	CollectionsCount int         `json:"collectionsCount" mapstructure:"collections_count"`
	ReSharesCount    int         `json:"reSharesCount" mapstructure:"reshares_count"`
	ID               string      `json:"id" mapstructure:"id"`
	Author           Author      `json:"author" mapstructure:"author"`
	IsCollected      bool        `json:"isCollected" mapstructure:"is_collected"`
	Subtype          string      `json:"subtype" mapstructure:"subtype"`
	Type             string      `json:"type" mapstructure:"type"`
	OwnerURI         string      `json:"ownerUri" mapstructure:"owner_uri"`
	Status           interface{} `json:"status" mapstructure:"status"`
	ReactionType     int         `json:"reactionType" mapstructure:"reaction_type"`
	Description      string      `json:"description" mapstructure:"description"`
	SharingURL       string      `json:"sharingUrl" mapstructure:"sharing_url"`
	URL              string      `json:"url" mapstructure:"url"`
	ReplyLimit       string      `json:"replyLimit" mapstructure:"reply_limit"`
	URI              string      `json:"uri" mapstructure:"uri"`
	LikerCount       int         `json:"likerCount" mapstructure:"likers_count"`
	ReactionsCount   int         `json:"reactionsCount" mapstructure:"reactions_count"`
	CommentsCount    int         `json:"commentsCount" mapstructure:"comments_count"`
	Position         int         `json:"position" mapstructure:"position"`
}
