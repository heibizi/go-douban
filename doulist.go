package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) DouList(id string) (DouListResult, error) {
	u, _ := url.JoinPath(apiUrl, "/doulist/", id)
	var o DouListResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 doulist 失败: %v", err)
	}
	return o, nil
}

type DouListResult struct {
	IsFollow           bool     `json:"isFollow" mapstructure:"is_follow"`
	ScreenshotTitle    string   `json:"screenshotTitle" mapstructure:"screenshot_title"`
	PlayableCount      int      `json:"playableCount" mapstructure:"playable_count"`
	ScreenshotURL      string   `json:"screenshotURL" mapstructure:"screenshot_url"`
	CreateTime         string   `json:"createTime" mapstructure:"create_time"`
	Owner              Owner    `json:"owner" mapstructure:"owner"`
	ScreenshotType     string   `json:"screenshotType" mapstructure:"screenshot_type"`
	ID                 string   `json:"id" mapstructure:"id"`
	Category           string   `json:"category" mapstructure:"category"`
	IsMergedCover      bool     `json:"isMergedCover" mapstructure:"is_merged_cover"`
	Title              string   `json:"title" mapstructure:"title"`
	IsSubjectSelection bool     `json:"isSubjectSelection" mapstructure:"is_subject_selection"`
	FollowersCount     int      `json:"followersCount" mapstructure:"followers_count"`
	IsPrivate          bool     `json:"isPrivate" mapstructure:"is_private"`
	SharingURL         string   `json:"sharingURL" mapstructure:"sharing_url"`
	Type               string   `json:"type" mapstructure:"type"`
	UpdateTime         string   `json:"updateTime" mapstructure:"update_time"`
	Tags               []string `json:"tags" mapstructure:"tags"`
	SyncingNote        *string  `json:"syncingNote" mapstructure:"syncing_note"`
	CoverURL           string   `json:"coverURL" mapstructure:"cover_url"`
	HeaderBGImage      string   `json:"headerBGImage" mapstructure:"header_bg_image"`
	DouListType        string   `json:"douListType" mapstructure:"doulist_type"`
	DoneCount          int      `json:"doneCount" mapstructure:"done_count"`
	Desc               string   `json:"desc" mapstructure:"desc"`
	FilterSwitch       struct {
		RatingRange bool `json:"ratingRange" mapstructure:"rating_range"`
	} `json:"filterSwitch" mapstructure:"filter_switch"`
	ItemsCount          int    `json:"itemsCount" mapstructure:"items_count"`
	WechatTimelineShare string `json:"wechatTimelineShare" mapstructure:"wechat_timeline_share"`
	URL                 string `json:"url" mapstructure:"url"`
	IsSysPrivate        bool   `json:"isSysPrivate" mapstructure:"is_sys_private"`
	URI                 string `json:"uri" mapstructure:"uri"`
	ListType            string `json:"listType" mapstructure:"list_type"`
}
