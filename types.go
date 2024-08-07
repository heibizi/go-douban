package douban

import "encoding/xml"

type (
	MediaType struct {
		Code string
		Name string
	}
	SearchSubjectsType struct {
		MediaType MediaType
		Tag       string
	}
	Option struct {
		Value string
		Name  string
	}
)

// web 相关
type (
	interestsRssResult struct {
		XMLName xml.Name `xml:"rss"`
		Items   []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"channel>item"`
	}
	// Detail 详情
	Detail struct {
		Intro      string  `mapstructure:"intro"`
		Cover      string  `mapstructure:"cover"`
		Rate       float64 `mapstructure:"rate"`
		Imdb       string  `mapstructure:"imdb"`
		Season     int     `mapstructure:"season"`
		EpisodeNum int     `mapstructure:"episode_num"`
		Title      string  `mapstructure:"title"`
		Year       int     `mapstructure:"year"`
	}
	InterestsRssInfo struct {
		Title string
		Url   string
		Date  string
		Type  InterestsType
	}
	NowPlaying struct {
		Id    string  `mapstructure:"id"`
		Title string  `mapstructure:"title"`
		Rate  float64 `mapstructure:"rate"`
		Cover string  `mapstructure:"cover"`
		Year  int     `mapstructure:"year"`
	}
	Later struct {
		Id    string `mapstructure:"id"`
		Title string `mapstructure:"title"`
		Cover string `mapstructure:"cover"`
		Url   string `mapstructure:"url"`
	}
	Top250 struct {
		Title string `mapstructure:"title"`
		Cover string `mapstructure:"cover"`
		Url   string `mapstructure:"url"`
	}
	Collect struct {
		Title string `mapstructure:"title"`
		Cover string `mapstructure:"cover"`
		Url   string `mapstructure:"url"`
	}
	Wish struct {
		Title string `mapstructure:"title"`
		Cover string `mapstructure:"cover"`
		Url   string `mapstructure:"url"`
		Date  string `mapstructure:"date"`
	}
	Do struct {
		Title string `mapstructure:"title"`
		Cover string `mapstructure:"cover"`
		Url   string `mapstructure:"url"`
	}
	UserInfo struct {
		Name string `mapstructure:"name"`
	}
	SearchSubject struct {
		EpisodesInfo string  `mapstructure:"episodes_info"`
		Rate         float64 `mapstructure:"rate"`
		CoverX       int     `mapstructure:"cover_x"`
		Title        string  `mapstructure:"title"`
		Url          string  `mapstructure:"url"`
		Playable     bool    `mapstructure:"playable"`
		Cover        string  `mapstructure:"cover"`
		Id           string  `mapstructure:"id"`
		CoverY       int     `mapstructure:"cover_y"`
		IsNew        bool    `mapstructure:"is_new"`
	}
)

// api 相关
type (
	ApiResult struct {
		Count int `json:"count" mapstructure:"count"`
		Start int `json:"start" mapstructure:"start"`
		Total int `json:"total" mapstructure:"total"`
	}
	Pic struct {
		Large  string `mapstructure:"large"`
		Normal string `mapstructure:"normal"`
	}
	BackgroundColorScheme struct {
		IsDark            bool   `mapstructure:"is_dark"`
		PrimaryColorLight string `mapstructure:"primary_color_light"`
		SecondaryColor    string `mapstructure:"secondary_color"`
		PrimaryColorDark  string `mapstructure:"primary_color_dark"`
	}
)
