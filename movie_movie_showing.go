package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) MovieMovieShowing(locId string, start int, count int, sort string) (MovieShowingResult, error) {
	u, _ := url.JoinPath(apiUrl, "/movie/movie_showing")
	var o MovieShowingResult
	params := map[string]string{"loc_id": locId, "start": strconv.Itoa(start), "count": strconv.Itoa(count), "sort": sort, "playable": "0"}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 movie_showing 失败: %v", err)
	}
	return o, nil
}

type MovieShowingResult struct {
	ApiResult   `mapstructure:",squash"`
	SharingInfo SharingInfo `json:"sharingInfo" mapstructure:"sharing_info"`
	Items       []MovieItem `json:"items" mapstructure:"items"`
}

type SharingInfo struct {
	SharingUrl          string `json:"sharingUrl" mapstructure:"sharing_url"`
	Title               string `json:"title" mapstructure:"title"`
	ScreenshotTitle     string `json:"screenshotTitle" mapstructure:"screenshot_title"`
	ScreenshotUrl       string `json:"screenshotUrl" mapstructure:"screenshot_url"`
	ScreenshotType      string `json:"screenshotType" mapstructure:"screenshot_type"`
	WechatTimelineShare string `json:"wechatTimelineShare" mapstructure:"wechat_timeline_share"`
}

type MovieItem struct {
	Comment       string      `json:"comment" mapstructure:"comment"`
	Rating        Rating      `json:"rating" mapstructure:"rating"`
	LineTicketURL string      `json:"lineTicketUrl" mapstructure:"lineticket_url"`
	HasLineWatch  bool        `json:"hasLineWatch" mapstructure:"has_linewatch"`
	Pic           Pic         `json:"pic" mapstructure:"pic"`
	HonorInfos    []HonorInfo `json:"honorInfos" mapstructure:"honor_infos"`
	URI           string      `json:"uri" mapstructure:"uri"`
	Photos        []string    `json:"photos" mapstructure:"photos"`
	VendorIcons   []string    `json:"vendorIcons" mapstructure:"vendor_icons"`
	Interest      string      `json:"interest" mapstructure:"interest"`
	Year          int         `json:"year" mapstructure:"year"`
	CardSubtitle  string      `json:"cardSubtitle" mapstructure:"card_subtitle"`
	Title         string      `json:"title" mapstructure:"title"`
	Type          string      `json:"type" mapstructure:"type"`
	ID            string      `json:"id" mapstructure:"id"`
	Tags          []Tag       `json:"tags" mapstructure:"tags"`
}
