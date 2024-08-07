package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) SubjectCollection(id string) (SubjectCollectionResult, error) {
	u, _ := url.JoinPath(apiUrl, "/subject_collection/", id)
	var o SubjectCollectionResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 subject_collection 失败: %v", err)
	}
	return o, nil
}

type SkyNetModel struct {
	VideoCount int    `json:"videoCount" mapstructure:"video_count"`
	PlayURI    string `json:"playUri" mapstructure:"play_uri"`
}

type SubjectCollectionResult struct {
	SubjectType           string                `json:"subjectType" mapstructure:"subject_type"`
	Subtitle              string                `json:"subtitle" mapstructure:"subtitle"`
	BackgroundColorScheme BackgroundColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
	IsFollow              bool                  `json:"isFollow" mapstructure:"is_follow"`
	UpdatedAt             string                `json:"updatedAt" mapstructure:"updated_at"`
	ScreenshotTitle       string                `json:"screenshotTitle" mapstructure:"screenshot_title"`
	ScreenshotURL         string                `json:"screenshotUrl" mapstructure:"screenshot_url"`
	SkyNetModel           SkyNetModel           `json:"skyNetModel" mapstructure:"skynet_model"`
	Total                 int                   `json:"total" mapstructure:"total"`
	ScreenshotType        string                `json:"screenshotType" mapstructure:"screenshot_type"`
	TypeIconBgText        string                `json:"typeIconBgText" mapstructure:"type_icon_bg_text"`
	Category              string                `json:"category" mapstructure:"category"`
	IsOfficial            bool                  `json:"isOfficial" mapstructure:"is_official"`
	IsMergedCover         bool                  `json:"isMergedCover" mapstructure:"is_merged_cover"`
	Title                 string                `json:"title" mapstructure:"title"`
	WxQRCode              string                `json:"wxQrCode" mapstructure:"wx_qr_code"`
	Personage             *string               `json:"personage" mapstructure:"personage"`
	ID                    string                `json:"id" mapstructure:"id"`
	FollowersCount        int                   `json:"followersCount" mapstructure:"followers_count"`
	ShowHeaderMask        bool                  `json:"showHeaderMask" mapstructure:"show_header_mask"`
	RelatedCharts         struct {
		Items []RankItem `json:"items" mapstructure:"items"`
		URI   string     `json:"uri" mapstructure:"uri"`
		Title string     `json:"title" mapstructure:"title"`
	} `json:"relatedCharts" mapstructure:"related_charts"`
	MediumName          string      `json:"mediumName" mapstructure:"medium_name"`
	Badge               Badge       `json:"badge" mapstructure:"badge"`
	RankType            string      `json:"rankType" mapstructure:"rank_type"`
	MergedTabs          []MergedTab `json:"mergedTabs" mapstructure:"merged_tabs"`
	Description         string      `json:"description" mapstructure:"description"`
	ShortName           string      `json:"shortName" mapstructure:"short_name"`
	NFollowers          int         `json:"nFollowers" mapstructure:"n_followers"`
	CoverURL            string      `json:"coverUrl" mapstructure:"cover_url"`
	HeaderBGImage       string      `json:"headerBgImage" mapstructure:"header_bg_image"`
	CanFollow           bool        `json:"canFollow" mapstructure:"can_follow"`
	ShowRank            bool        `json:"showRank" mapstructure:"show_rank"`
	ChartID             int         `json:"chartId" mapstructure:"chart_id"`
	CompleteAt          *string     `json:"completeAt" mapstructure:"complete_at"`
	Name                string      `json:"name" mapstructure:"name"`
	DoneCount           int         `json:"doneCount" mapstructure:"done_count"`
	SharingURL          string      `json:"sharingUrl" mapstructure:"sharing_url"`
	WxAppCode           string      `json:"wxAppCode" mapstructure:"wx_app_code"`
	SubjectCount        int         `json:"subjectCount" mapstructure:"subject_count"`
	ItemsCount          int         `json:"itemsCount" mapstructure:"items_count"`
	WechatTimelineShare string      `json:"wechatTimelineShare" mapstructure:"wechat_timeline_share"`
	CollectCount        int         `json:"collectCount" mapstructure:"collect_count"`
	URL                 string      `json:"url" mapstructure:"url"`
	Type                string      `json:"type" mapstructure:"type"`
	IsBadgeChart        bool        `json:"isBadgeChart" mapstructure:"is_badge_chart"`
	URI                 string      `json:"uri" mapstructure:"uri"`
	MiniProgramPage     string      `json:"miniProgramPage" mapstructure:"mini_program_page"`
	IconFgImage         string      `json:"iconFgImage" mapstructure:"icon_fg_image"`
	IconText            string      `json:"iconText" mapstructure:"icon_text"`
	MoreDescription     string      `json:"moreDescription" mapstructure:"more_description"`
	ListType            string      `json:"listType" mapstructure:"list_type"`
	MiniProgramName     string      `json:"miniProgramName" mapstructure:"mini_program_name"`
	FinishSoon          bool        `json:"finishSoon" mapstructure:"finish_soon"`
	Display             Display     `json:"display" mapstructure:"display"`
	User                *string     `json:"user" mapstructure:"user"`
}

type Badge struct {
	Status      int     `json:"status" mapstructure:"status"`
	MinDoneCnt  int     `json:"minDoneCnt" mapstructure:"min_done_cnt"`
	Target      *string `json:"target" mapstructure:"target"`
	Title       string  `json:"title" mapstructure:"title"`
	Series      string  `json:"series" mapstructure:"series"`
	URI         string  `json:"uri" mapstructure:"uri"`
	IsNew       bool    `json:"isNew" mapstructure:"is_new"`
	Chart       *string `json:"chart" mapstructure:"chart"`
	BgColor     string  `json:"bgColor" mapstructure:"bg_color"`
	UserBadgeID *string `json:"userBadgeId" mapstructure:"user_badge_id"`
	ReceivedAt  *string `json:"receivedAt" mapstructure:"received_at"`
	ID          int     `json:"id" mapstructure:"id"`
	Icon        Icon    `json:"icon" mapstructure:"icon"`
}

type MergedTab struct {
	Category   string           `json:"category" mapstructure:"category"`
	Items      []MergedTabItem  `json:"items" mapstructure:"items"`
	GroupItems []MergedTabGroup `json:"groupItems" mapstructure:"group_items"`
}

type MergedTabItem struct {
	Current bool   `json:"current" mapstructure:"current"`
	ID      string `json:"id" mapstructure:"id"`
	Name    string `json:"name" mapstructure:"name"`
	URI     string `json:"uri" mapstructure:"uri"`
}

type MergedTabGroup struct {
	Category string          `json:"category" mapstructure:"category"`
	Items    []MergedTabItem `json:"items" mapstructure:"items"`
}

type RelatedChartItem struct {
	RankType              string                `json:"rankType" mapstructure:"rank_type"`
	ItemsCount            int                   `json:"itemsCount" mapstructure:"items_count"`
	IsMergedCover         bool                  `json:"isMergedCover" mapstructure:"is_merged_cover"`
	BackgroundColorScheme BackgroundColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
	ShortName             string                `json:"shortName" mapstructure:"short_name"`
	URL                   string                `json:"url" mapstructure:"url"`
	Name                  string                `json:"name" mapstructure:"name"`
	NFollowers            int                   `json:"nFollowers" mapstructure:"n_followers"`
	CoverURL              string                `json:"coverUrl" mapstructure:"cover_url"`
	URI                   string                `json:"uri" mapstructure:"uri"`
	IconFgImage           string                `json:"iconFgImage" mapstructure:"icon_fg_image"`
	SubjectType           string                `json:"subjectType" mapstructure:"subject_type"`
	ShowHeaderMask        bool                  `json:"showHeaderMask" mapstructure:"show_header_mask"`
	ListType              string                `json:"listType" mapstructure:"list_type"`
	ID                    string                `json:"id" mapstructure:"id"`
	Subtitle              string                `json:"subtitle" mapstructure:"subtitle"`
	MediumName            string                `json:"mediumName" mapstructure:"medium_name"`
	Type                  string                `json:"type" mapstructure:"type"`
	Subjects              []Subject             `json:"subjects" mapstructure:"subjects"`
	DoneCount             int                   `json:"doneCount" mapstructure:"done_count"`
	SharingURL            string                `json:"sharingUrl" mapstructure:"sharing_url"`
}

type Subject struct {
	Rating            Rating      `json:"rating" mapstructure:"rating"`
	ControversyReason string      `json:"controversyReason" mapstructure:"controversy_reason"`
	Title             string      `json:"title" mapstructure:"title"`
	URL               string      `json:"url" mapstructure:"url"`
	Pic               Pic         `json:"pic" mapstructure:"pic"`
	URI               string      `json:"uri" mapstructure:"uri"`
	CoverURL          string      `json:"coverUrl" mapstructure:"cover_url"`
	IsReleased        bool        `json:"isReleased" mapstructure:"is_released"`
	IsShow            bool        `json:"isShow" mapstructure:"is_show"`
	IsPlayable        bool        `json:"isPlayable" mapstructure:"is_playable"`
	Subtype           string      `json:"subtype" mapstructure:"subtype"`
	CardSubtitle      string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	ColorScheme       ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	Type              string      `json:"type" mapstructure:"type"`
	ID                string      `json:"id" mapstructure:"id"`
	NullRatingReason  string      `json:"nullRatingReason" mapstructure:"null_rating_reason"`
	SharingURL        string      `json:"sharingUrl" mapstructure:"sharing_url"`
}
