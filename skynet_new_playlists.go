package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

// SkyNetNewPlayLists 影视片单
func (c *ApiClient) SkyNetNewPlayLists(category string, mediaType MediaType, start int, count int) (SkyNetNewPlayListsResult, error) {
	u, _ := url.JoinPath(apiUrl, "/skynet/new_playlists")
	var o SkyNetNewPlayListsResult
	params := map[string]string{"category": category, "subject_type": mediaType.Code, "start": strconv.Itoa(start), "count": strconv.Itoa(count)}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 sky_net/new_playlists 失败: %v", err)
	}
	return o, nil
}

type SkyNetNewPlayListsResult struct {
	ApiResult  `mapstructure:",squash"`
	SharingURL string `json:"sharingUrl" mapstructure:"sharing_url"`
	Data       []struct {
		Category string                   `json:"category" mapstructure:"category"`
		Total    int                      `json:"total" mapstructure:"total"`
		Items    []SkyNetNewPlayListsItem `json:"items" mapstructure:"items"`
	} `json:"data" mapstructure:"data"`
}

type SkyNetNewPlayListsItem struct {
	BackgroundColorScheme BackgroundColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
	Total                 int                   `json:"total" mapstructure:"total"`
	ID                    string                `json:"id" mapstructure:"id"`
	Category              string                `json:"category" mapstructure:"category"`
	IsMergedCover         bool                  `json:"isMergedCover" mapstructure:"is_merged_cover"`
	Title                 string                `json:"title" mapstructure:"title"`
	IconText              string                `json:"iconText" mapstructure:"icon_text"`
	FollowersCount        int                   `json:"followersCount" mapstructure:"followers_count"`
	Type                  string                `json:"type" mapstructure:"type"`
	RankType              string                `json:"rankType" mapstructure:"rank_type"`
	CoverURL              string                `json:"coverUrl" mapstructure:"cover_url"`
	HeaderBgImage         string                `json:"headerBgImage" mapstructure:"header_bg_image"`
	DoneCount             int                   `json:"doneCount" mapstructure:"done_count"`
	SubjectCount          int                   `json:"subjectCount" mapstructure:"subject_count"`
	ItemsCount            int                   `json:"itemsCount" mapstructure:"items_count"`
	SharingURL            string                `json:"sharingUrl" mapstructure:"sharing_url"`
	CollectCount          int                   `json:"collectCount" mapstructure:"collect_count"`
	URL                   string                `json:"url" mapstructure:"url"`
	IsBadgeChart          bool                  `json:"isBadgeChart" mapstructure:"is_badge_chart"`
	URI                   string                `json:"uri" mapstructure:"uri"`
	ItemType              string                `json:"itemType" mapstructure:"item_type"`
	FinishSoon            bool                  `json:"finishSoon" mapstructure:"finish_soon"`
	ListType              string                `json:"listType" mapstructure:"list_type"`
}
