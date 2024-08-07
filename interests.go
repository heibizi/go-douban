package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Interests(mediaType MediaType, id string) (InterestsResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/interests")
	var o InterestsResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 interests 失败: %v", err)
	}
	return o, nil
}

type InterestsResult struct {
	ApiResult           `mapstructure:",squash"`
	WechatTimelineShare string     `json:"wechatTimelineShare" mapstructure:"wechatTimelineShare"`
	Interests           []Interest `json:"interests" mapstructure:"interests"`
}

type User struct {
	Loc                Location `json:"loc" mapstructure:"loc"`
	RegTime            string   `json:"regTime" mapstructure:"reg_time"`
	Followed           bool     `json:"followed" mapstructure:"followed"`
	Name               string   `json:"name" mapstructure:"name"`
	UID                string   `json:"uid" mapstructure:"uid"`
	URL                string   `json:"url" mapstructure:"url"`
	Gender             string   `json:"gender" mapstructure:"gender"`
	URI                string   `json:"uri" mapstructure:"uri"`
	Kind               string   `json:"kind" mapstructure:"kind"`
	AvatarSideIcon     string   `json:"avatarSideIcon" mapstructure:"avatar_side_icon"`
	IsClub             bool     `json:"isClub" mapstructure:"is_club"`
	Remark             string   `json:"remark" mapstructure:"remark"`
	AvatarSideIconID   string   `json:"avatarSideIconID" mapstructure:"avatar_side_icon_id"`
	InBlacklist        bool     `json:"inBlacklist" mapstructure:"in_blacklist"`
	AvatarSideIconType int      `json:"avatarSideIconType" mapstructure:"avatar_side_icon_type"`
	Type               string   `json:"type" mapstructure:"type"`
	ID                 string   `json:"id" mapstructure:"id"`
	Avatar             string   `json:"avatar" mapstructure:"avatar"`
}

type Interest struct {
	Comment             string   `json:"comment" mapstructure:"comment"`
	Rating              Rating   `json:"rating" mapstructure:"rating"`
	SharingURL          string   `json:"sharingUrl" mapstructure:"sharing_url"`
	ShowTimeTip         bool     `json:"showTimeTip" mapstructure:"show_time_tip"`
	IsVoted             bool     `json:"isVoted" mapstructure:"is_voted"`
	URI                 string   `json:"uri" mapstructure:"uri"`
	Platforms           []string `json:"platforms" mapstructure:"platforms"`
	VoteCount           int      `json:"voteCount" mapstructure:"vote_count"`
	CreateTime          string   `json:"createTime" mapstructure:"create_time"`
	Status              string   `json:"status" mapstructure:"status"`
	User                User     `json:"user" mapstructure:"user"`
	IPLocation          string   `json:"ipLocation" mapstructure:"ip_location"`
	RecommendReason     string   `json:"recommendReason" mapstructure:"recommend_reason"`
	UserDoneDesc        string   `json:"userDoneDesc" mapstructure:"user_done_desc"`
	ID                  string   `json:"id" mapstructure:"id"`
	WeChatTimelineShare string   `json:"wechatTimelineShare" mapstructure:"wechat_timeline_share"`
}
