package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Detail(mediaType MediaType, id string) (DetailResult, error) {
	u, _ := url.JoinPath(apiUrl, "/"+mediaType.Code+"/", id)
	var o DetailResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 detail 失败: %v", err)
	}
	return o, nil
}

type DetailResult struct {
	Rating            Rating   `json:"rating" mapstructure:"rating"`
	LineTicketURL     string   `json:"lineTicketUrl" mapstructure:"lineticket_url"`
	ControversyReason string   `json:"controversyReason" mapstructure:"controversy_reason"`
	PubDate           []string `json:"pubDate" mapstructure:"pubdate"`
	//LastEpisodeNumber          *int             `json:"lastEpisodeNumber" mapstructure:"last_episode_number"`
	//InterestControlInfo        *interface{}     `json:"interestControlInfo" mapstructure:"interest_control_info"`
	Pic             Pic     `json:"pic" mapstructure:"pic"`
	Year            int     `json:"year" mapstructure:"year"`
	VendorCount     int     `json:"vendorCount" mapstructure:"vendor_count"`
	BodyBgColor     string  `json:"bodyBgColor" mapstructure:"body_bg_color"`
	IsTv            bool    `json:"isTv" mapstructure:"is_tv"`
	CardSubtitle    string  `json:"cardSubtitle" mapstructure:"card_subtitle"`
	AlbumNoInteract bool    `json:"albumNoInteract" mapstructure:"album_no_interact"`
	TicketPriceInfo string  `json:"ticketPriceInfo" mapstructure:"ticket_price_info"`
	PrePlayableDate *string `json:"prePlayableDate" mapstructure:"pre_playable_date"`
	CanRate         bool    `json:"canRate" mapstructure:"can_rate"`
	//HeadInfo                   *interface{}     `json:"headInfo" mapstructure:"head_info"`
	//ForumInfo                  *interface{}     `json:"forumInfo" mapstructure:"forum_info"`
	//ShareActivities            []interface{}    `json:"shareActivities" mapstructure:"share_activities"`
	//Webisode                   *interface{}     `json:"webisode" mapstructure:"webisode"`
	ID                         string           `json:"id" mapstructure:"id"`
	GalleryTopicCount          int              `json:"galleryTopicCount" mapstructure:"gallery_topic_count"`
	Languages                  []string         `json:"languages" mapstructure:"languages"`
	Genres                     []string         `json:"genres" mapstructure:"genres"`
	ReviewCount                int              `json:"reviewCount" mapstructure:"review_count"`
	VariableModules            []VariableModule `json:"variableModules" mapstructure:"variable_modules"`
	Title                      string           `json:"title" mapstructure:"title"`
	Intro                      string           `json:"intro" mapstructure:"intro"`
	InterestCmtEarlierTipTitle string           `json:"interestCmtEarlierTipTitle" mapstructure:"interest_cmt_earlier_tip_title"`
	HasLineWatch               bool             `json:"hasLineWatch" mapstructure:"has_linewatch"`
	CommentCount               int              `json:"commentCount" mapstructure:"comment_count"`
	ForumTopicCount            int              `json:"forumTopicCount" mapstructure:"forum_topic_count"`
	TicketPromoText            string           `json:"ticketPromoText" mapstructure:"ticket_promo_text"`
	//WebviewInfo                *interface{}     `json:"webviewInfo" mapstructure:"webview_info"`
	IsReleased bool `json:"isReleased" mapstructure:"is_released"`
	//Vendors                    []interface{}    `json:"vendors" mapstructure:"vendors"`
	Actors []Actor `json:"actors" mapstructure:"actors"`
	//Interest                   *interface{}     `json:"interest" mapstructure:"interest"`
	Subtype       string      `json:"subtype" mapstructure:"subtype"`
	EpisodesCount int         `json:"episodesCount" mapstructure:"episodes_count"`
	ColorScheme   ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	Type          string      `json:"type" mapstructure:"type"`
	//LineWatches
	InfoURL           string      `json:"infoUrl" mapstructure:"info_url"`
	Tags              []string    `json:"tags" mapstructure:"tags"`
	VendorDesc        string      `json:"vendorDesc" mapstructure:"vendor_desc"`
	Durations         []string    `json:"durations" mapstructure:"durations"`
	Cover             DetailCover `json:"cover" mapstructure:"cover"`
	CoverURL          string      `json:"coverUrl" mapstructure:"cover_url"`
	Trailers          []Trailer   `json:"trailer" mapstructure:"trailer"`
	HeaderBgColor     string      `json:"headerBgColor" mapstructure:"header_bg_color"`
	IsDoubanIntro     bool        `json:"isDoubanIntro" mapstructure:"is_douban_intro"`
	TicketVendorIcons []string    `json:"ticketVendorIcons" mapstructure:"ticket_vendor_icons"`
	HonorInfos        []HonorInfo `json:"honorInfos" mapstructure:"honor_infos"`
	SharingURL        string      `json:"sharingUrl" mapstructure:"sharing_url"`
	//SubjectCollections
	WechatTimelineShare string   `json:"wechatTimelineShare" mapstructure:"wechat_timeline_share"`
	RestrictiveIconUrl  string   `json:"restrictiveIconUrl" mapstructure:"restrictive_icon_url"`
	RateInfo            string   `json:"rateInfo" mapstructure:"rate_info"`
	ReleaseDate         string   `json:"releaseDate" mapstructure:"release_date"`
	Countries           []string `json:"countries" mapstructure:"countries"`
	OriginalTitle       string   `json:"originalTitle" mapstructure:"original_title"`
	URI                 string   `json:"uri" mapstructure:"uri"`
	WebisodeCount       int      `json:"webisodeCount" mapstructure:"webisode_count"`
	EpisodesInfo        string   `json:"episodesInfo" mapstructure:"episodes_info"`
	URL                 string   `json:"url" mapstructure:"url"`
	Directors           []struct {
		Name string `json:"name" mapstructure:"name"`
	} `json:"directors" mapstructure:"directors"`
	IsShow           bool     `json:"isShow" mapstructure:"is_show"`
	VendorIcons      []string `json:"vendorIcons" mapstructure:"vendor_icons"`
	PreReleaseDesc   string   `json:"preReleaseDesc" mapstructure:"pre_release_desc"`
	Video            string   `json:"video" mapstructure:"video"`
	Aka              []string `json:"aka" mapstructure:"aka"`
	IsRestrictive    bool     `json:"isRestrictive" mapstructure:"is_restrictive"`
	NullRatingReason string   `json:"nullRatingReason" mapstructure:"null_rating_reason"`
}

type VariableModule struct {
	SubModules []SubModule `json:"subModules" mapstructure:"sub_modules"`
	ID         string      `json:"id" mapstructure:"id"`
}

type SubModule struct {
	ID       string        `json:"id" mapstructure:"id"`
	Data     *interface{}  `json:"data,omitempty" mapstructure:"data"`
	DataType *string       `json:"dataType,omitempty" mapstructure:"data_type"`
	SortBy   *[]SortOption `json:"sortBy,omitempty" mapstructure:"sort_by"`
}

type SortOption struct {
	Name string `json:"name" mapstructure:"name"`
	ID   string `json:"id" mapstructure:"id"`
}

type DetailCover struct {
	Description string `json:"description" mapstructure:"description"`
	Author      Author `json:"author" mapstructure:"author"`
	URL         string `json:"url" mapstructure:"url"`
	Image       Image  `json:"image" mapstructure:"image"`
	URI         string `json:"uri" mapstructure:"uri"`
	CreateTime  string `json:"createTime" mapstructure:"create_time"`
	Position    int    `json:"position" mapstructure:"position"`
	OwnerURI    string `json:"ownerUri" mapstructure:"owner_uri"`
	Type        string `json:"type" mapstructure:"type"`
	ID          string `json:"id" mapstructure:"id"`
	SharingURL  string `json:"sharingUrl" mapstructure:"sharing_url"`
}

type Author struct {
	Loc                Location `json:"loc" mapstructure:"loc"`
	Kind               string   `json:"kind" mapstructure:"kind"`
	Name               string   `json:"name" mapstructure:"name"`
	RegYear            int      `json:"regYear" mapstructure:"reg_year"`
	AvatarSideIconType int      `json:"avatarSideIconType" mapstructure:"avatar_side_icon_type"`
	URL                string   `json:"url" mapstructure:"url"`
	ID                 string   `json:"id" mapstructure:"id"`
	RegTime            string   `json:"regTime" mapstructure:"reg_time"`
	URI                string   `json:"uri" mapstructure:"uri"`
	Avatar             string   `json:"avatar" mapstructure:"avatar"`
	IsClub             bool     `json:"isClub" mapstructure:"is_club"`
	Type               string   `json:"type" mapstructure:"type"`
	AvatarSideIcon     string   `json:"avatarSideIcon" mapstructure:"avatar_side_icon"`
	UID                string   `json:"uid" mapstructure:"uid"`
}

type Image struct {
	Normal       ImageSize  `json:"normal" mapstructure:"normal"`
	Large        ImageSize  `json:"large" mapstructure:"large"`
	Raw          *ImageSize `json:"raw,omitempty" mapstructure:"raw"`
	Small        ImageSize  `json:"small" mapstructure:"small"`
	PrimaryColor string     `json:"primaryColor" mapstructure:"primary_color"`
	IsAnimated   bool       `json:"isAnimated" mapstructure:"is_animated"`
}

type ImageSize struct {
	URL    string `json:"url" mapstructure:"url"`
	Width  int    `json:"width" mapstructure:"width"`
	Height int    `json:"height" mapstructure:"height"`
	Size   int    `json:"size" mapstructure:"size"`
}
