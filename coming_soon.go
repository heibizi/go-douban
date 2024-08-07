package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) ComingSoon(mediaType MediaType, start int, count int, sort string, area string) (ComingSoonResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/coming_soon")
	var o ComingSoonResult
	params := map[string]string{"start": strconv.Itoa(start), "count": strconv.Itoa(count), "sortby": sort, "area": area}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 coming_soon 失败: %v", err)
	}
	return o, nil
}

type ComingSoonResult struct {
	ApiResult `mapstructure:",squash"`
	Subjects  []ComingSoonSubject `json:"subjects" mapstructure:"subjects"`
	Filters   []struct {
		Items []FilterItem `json:"items" mapstructure:"items"`
		Type  string       `json:"type" mapstructure:"type"`
		Title string       `json:"title" mapstructure:"title"`
	}
}
type FilterItem struct {
	Tag     string `json:"tag" mapstructure:"tag"`
	Checked bool   `json:"checked" mapstructure:"checked"`
	Id      string `json:"id" mapstructure:"id"`
}

type Highlight struct {
	Kind string `json:"kind" mapstructure:"kind"`
	Desc string `json:"desc" mapstructure:"desc"`
}

type Trailer struct {
	SharingURL string `json:"sharingUrl" mapstructure:"sharing_url"`
	VideoURL   string `json:"videoUrl" mapstructure:"video_url"`
	Title      string `json:"title" mapstructure:"title"`
	TypeName   string `json:"typeName" mapstructure:"type_name"`
	URI        string `json:"uri" mapstructure:"uri"`
	CoverURL   string `json:"coverUrl" mapstructure:"cover_url"`
	TermNum    int    `json:"termNum" mapstructure:"term_num"`
	NComments  int    `json:"nComments" mapstructure:"n_comments"`
	CreateTime string `json:"createTime" mapstructure:"create_time"`
	FileSize   int    `json:"fileSize" mapstructure:"file_size"`
	Runtime    string `json:"runtime" mapstructure:"runtime"`
	Type       string `json:"type" mapstructure:"type"`
	ID         string `json:"id" mapstructure:"id"`
	Desc       string `json:"desc" mapstructure:"desc"`
}

type ComingSoonSubject struct {
	Rating            Rating      `json:"rating" mapstructure:"rating"`
	ControversyReason string      `json:"controversyReason" mapstructure:"controversy_reason"`
	PubDate           []string    `json:"pubDate" mapstructure:"pubdate"`
	WishCount         int         `json:"wishCount" mapstructure:"wish_count"`
	Pic               Pic         `json:"pic" mapstructure:"pic"`
	IsShow            bool        `json:"isShow" mapstructure:"is_show"`
	VendorIcons       []string    `json:"vendorIcons" mapstructure:"vendor_icons"`
	Year              string      `json:"year" mapstructure:"year"`
	CardSubtitle      string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	ID                string      `json:"id" mapstructure:"id"`
	Genres            []string    `json:"genres" mapstructure:"genres"`
	Title             string      `json:"title" mapstructure:"title"`
	Highlights        []Highlight `json:"highlights" mapstructure:"highlights"`
	IsReleased        bool        `json:"isReleased" mapstructure:"is_released"`
	Actors            []Actor     `json:"actors" mapstructure:"actors"`
	Interest          interface{} `json:"interest" mapstructure:"interest"`
	ColorScheme       ColorScheme `json:"colorScheme" mapstructure:"color_scheme"`
	Type              string      `json:"type" mapstructure:"type"`
	HasLinewatch      bool        `json:"hasLinewatch" mapstructure:"has_linewatch"`
	CoverURL          string      `json:"coverUrl" mapstructure:"cover_url"`
	Photos            []string    `json:"photos" mapstructure:"photos"`
	SharingURL        string      `json:"sharingUrl" mapstructure:"sharing_url"`
	URL               string      `json:"url" mapstructure:"url"`
	ReleaseDate       string      `json:"releaseDate" mapstructure:"release_date"`
	URI               string      `json:"uri" mapstructure:"uri"`
	Subtype           string      `json:"subtype" mapstructure:"subtype"`
	Directors         []struct {
		Name string `json:"name" mapstructure:"name"`
	} `json:"directors" mapstructure:"directors"`
	Intro            string  `json:"intro" mapstructure:"intro"`
	NullRatingReason string  `json:"nullRatingReason" mapstructure:"null_rating_reason"`
	AlbumNoInteract  bool    `json:"albumNoInteract" mapstructure:"album_no_interact"`
	Trailer          Trailer `json:"trailer" mapstructure:"trailer"`
}
