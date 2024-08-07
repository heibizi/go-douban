package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) RankList(mediaType MediaType, start int, count int) (RankList, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/rank_list")
	var o RankList
	params := map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 rank_list 失败: %v", err)
	}
	return o, nil
}

type RankList struct {
	SharingURL string `json:"sharingUrl" mapstructure:"sharing_url"`
	Groups     []struct {
		Title               string               `json:"title" mapstructure:"title"`
		Type                string               `json:"type" mapstructure:"type"`
		SelectedCollections []SelectedCollection `json:"selectedCollections" mapstructure:"selected_collections"`
		Uri                 string               `json:"uri" mapstructure:"uri"`
		Tabs                []Tab                `json:"tabs" mapstructure:"tabs,omitempty"`
	} `json:"groups" mapstructure:"groups"`
	Title string `json:"title" mapstructure:"title"`
}

type SelectedCollection struct {
	SubjectType           string      `json:"subjectType" mapstructure:"subject_type"`
	Subtitle              string      `json:"subtitle" mapstructure:"subtitle"`
	BackgroundColorScheme ColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
	IsFollow              bool        `json:"isFollow" mapstructure:"is_follow"`
	UpdatedAt             string      `json:"updatedAt" mapstructure:"updated_at"`
	Rank                  bool        `json:"rank" mapstructure:"rank"`
	SkyNetModel           *string     `json:"skyNetModel" mapstructure:"skynet_model"`
	Total                 int         `json:"total" mapstructure:"total"`
	TypeIconBgText        string      `json:"typeIconBgText" mapstructure:"type_icon_bg_text"`
	Category              string      `json:"category" mapstructure:"category"`
	IsOfficial            bool        `json:"isOfficial" mapstructure:"is_official"`
	IsMergedCover         bool        `json:"isMergedCover" mapstructure:"is_merged_cover"`
	HeaderFgImage         string      `json:"headerFgImage" mapstructure:"header_fg_image"`
	CollectionTitle       string      `json:"title" mapstructure:"title"`
	WxQRCode              string      `json:"wxQrCode" mapstructure:"wx_qr_code"`
	FinishSoon            bool        `json:"finishSoon" mapstructure:"finish_soon"`
	Id                    string      `json:"id" mapstructure:"id"`
	FollowersCount        int         `json:"followersCount" mapstructure:"followers_count"`
	ShowHeaderMask        bool        `json:"showHeaderMask" mapstructure:"show_header_mask"`
	MediumName            string      `json:"mediumName" mapstructure:"medium_name"`
	CollectionType        string      `json:"type" mapstructure:"type"`
	RankType              string      `json:"rankType" mapstructure:"rank_type"`
	Description           string      `json:"description" mapstructure:"description"`
	ShortName             string      `json:"shortName" mapstructure:"short_name"`
	NFollowers            int         `json:"nFollowers" mapstructure:"n_followers"`
	CoverURL              string      `json:"coverUrl" mapstructure:"cover_url"`
	HeaderBgImage         string      `json:"headerBgImage" mapstructure:"header_bg_image"`
	ShowRank              bool        `json:"showRank" mapstructure:"show_rank"`
	ChartID               int         `json:"chartId" mapstructure:"chart_id"`
	DoneCount             int         `json:"doneCount" mapstructure:"done_count"`
	Name                  string      `json:"name" mapstructure:"name"`
	WxAppCode             string      `json:"wxAppCode" mapstructure:"wx_app_code"`
	SubjectCount          int         `json:"subjectCount" mapstructure:"subject_count"`
	ItemsCount            int         `json:"itemsCount" mapstructure:"items_count"`
	SharingURL            string      `json:"sharingUrl" mapstructure:"sharing_url"`
	CollectCount          int         `json:"collectCount" mapstructure:"collect_count"`
	URL                   string      `json:"url" mapstructure:"url"`
	Items                 []RankItem  `json:"items" mapstructure:"items"`
	IsBadgeChart          bool        `json:"isBadgeChart" mapstructure:"is_badge_chart"`
	CollectionUri         string      `json:"uri" mapstructure:"uri"`
	MiniProgramPage       string      `json:"miniProgramPage" mapstructure:"mini_program_page"`
	IconFgImage           string      `json:"iconFgImage" mapstructure:"icon_fg_image"`
	IconText              string      `json:"iconText" mapstructure:"icon_text"`
	MoreDescription       string      `json:"moreDescription" mapstructure:"more_description"`
	ListType              string      `json:"listType" mapstructure:"list_type"`
	MiniProgramName       string      `json:"miniProgramName" mapstructure:"mini_program_name"`
	Display               Display     `json:"display" mapstructure:"display"`
	User                  *string     `json:"user" mapstructure:"user"`
}

type ColorScheme struct {
	IsDark            bool      `json:"isDark" mapstructure:"is_dark"`
	PrimaryColorLight string    `json:"primaryColorLight" mapstructure:"primary_color_light"`
	SecondaryColor    string    `json:"secondaryColor" mapstructure:"secondary_color"`
	PrimaryColorDark  string    `json:"primaryColorDark" mapstructure:"primary_color_dark"`
	BaseColor         []float64 `json:"baseColor" mapstructure:"_base_color,omitempty"`
	AvgColor          []float64 `json:"avgColor" mapstructure:"_avg_color,omitempty"`
}

type RankItem struct {
	Rating            Rating      `json:"rating" mapstructure:"rating"`
	ControversyReason string      `json:"controversyReason" mapstructure:"controversy_reason"`
	PubDate           []string    `json:"pubDate" mapstructure:"pubdate"`
	RankValue         int         `json:"rankValue" mapstructure:"rank_value"`
	Pic               Picture     `json:"pic" mapstructure:"pic"`
	HonorInfos        []HonorInfo `json:"honorInfos" mapstructure:"honor_infos"`
	IsShow            bool        `json:"isShow" mapstructure:"is_show"`
	Year              string      `json:"year" mapstructure:"year"`
	CardSubtitle      string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	Id                string      `json:"id" mapstructure:"id"`
	Genres            []string    `json:"genres" mapstructure:"genres"`
	TrendDown         bool        `json:"trendDown" mapstructure:"trend_down"`
	Title             string      `json:"title" mapstructure:"title"`
	TrendEqual        bool        `json:"trendEqual" mapstructure:"trend_equal"`
	TrendUp           bool        `json:"trendUp" mapstructure:"trend_up"`
	IsReleased        bool        `json:"isReleased" mapstructure:"is_released"`
	RankValueChanged  int         `json:"rankValueChanged" mapstructure:"rank_value_changed"`
	Actors            []Actor     `json:"actors" mapstructure:"actors"`
	ColorScheme       ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	ItemType          string      `json:"type" mapstructure:"type"`
	HasLineWatch      bool        `json:"hasLineWatch" mapstructure:"has_linewatch"`
	CoverURL          string      `json:"coverUrl" mapstructure:"cover_url"`
	SharingURL        string      `json:"sharingUrl" mapstructure:"sharing_url"`
	ItemURL           string      `json:"url" mapstructure:"url"`
	ReleaseDate       *string     `json:"releaseDate" mapstructure:"release_date"`
	URI               string      `json:"uri" mapstructure:"uri"`
	SubType           string      `json:"subtype" mapstructure:"subtype"`
	Directors         []Actor     `json:"directors" mapstructure:"directors"`
	Intro             string      `json:"intro" mapstructure:"intro"`
	AlbumNoInteract   bool        `json:"albumNoInteract" mapstructure:"album_no_interact"`
	NullRatingReason  string      `json:"nullRatingReason" mapstructure:"null_rating_reason"`
}

type Picture struct {
	Large  string `json:"large" mapstructure:"large"`
	Normal string `json:"normal" mapstructure:"normal"`
}

type Rating struct {
	Count     int     `json:"count" mapstructure:"count"`
	Max       int     `json:"max" mapstructure:"max"`
	StarCount float64 `json:"starCount" mapstructure:"star_count"`
	Value     float64 `json:"value" mapstructure:"value"`
}

type Display struct {
	Layout string `json:"layout" mapstructure:"layout"`
}

type Tab struct {
	Key   string `json:"key" mapstructure:"key"`
	Title string `json:"title" mapstructure:"title"`
}
