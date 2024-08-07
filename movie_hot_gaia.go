package douban

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *ApiClient) MovieHotGaia(locId string, start int, count int, sort string) (HotGaiaResult, error) {
	u, _ := url.JoinPath(apiUrl, "/movie/hot_gaia")
	var o HotGaiaResult
	params := map[string]string{"loc_id": locId, "start": strconv.Itoa(start), "count": strconv.Itoa(count), "sort": sort, "playable": "0"}
	err := c.get(u, params, &o)
	if err != nil {
		return o, fmt.Errorf("获取 hot_gaia 失败: %v", err)
	}
	return o, nil
}

type HotGaiaResult struct {
	ApiResult   `mapstructure:",squash"`
	SharingInfo SharingInfo `json:"sharingInfo" mapstructure:"sharing_info"`
	Items       []MovieItem `json:"items" mapstructure:"items"`
	Areas       []string    `json:"areas" mapstructure:"areas"`
}
