package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) HotCities() (HotCitiesResult, error) {
	u, _ := url.JoinPath(apiUrl, "/hot_cities")
	var o HotCitiesResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 hot_cities 失败: %v", err)
	}
	return o, nil
}

type HotCitiesResult struct {
	Locations []Location `json:"locations" mapstructure:"locations"`
}

type Location struct {
	UID  string `json:"uid" mapstructure:"uid"`
	Name string `json:"name" mapstructure:"name"`
	Id   string `json:"id" mapstructure:"id"`
}
