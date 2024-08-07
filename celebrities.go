package douban

import (
	"fmt"
	"net/url"
)

func (c *ApiClient) Celebrities(mediaType MediaType, id string) (CelebritiesResult, error) {
	u, _ := url.JoinPath(apiUrl, "/", mediaType.Code, "/", id, "/celebrities")
	var o CelebritiesResult
	err := c.get(u, nil, &o)
	if err != nil {
		return o, fmt.Errorf("获取 celebrities 失败: %v", err)
	}
	return o, nil
}

type CelebritiesResult struct {
	Total     int        `json:"total" mapstructure:"total"`
	Directors []Director `json:"directors" mapstructure:"directors"`
	Actors    []Actor    `json:"actors" mapstructure:"actors"`
}

type Avatar struct {
	Large  string `json:"large" mapstructure:"large"`
	Normal string `json:"normal" mapstructure:"normal"`
}

type Director struct {
	Name            string   `json:"name" mapstructure:"name"`
	Roles           []string `json:"roles" mapstructure:"roles"`
	Title           string   `json:"title" mapstructure:"title"`
	URL             string   `json:"url" mapstructure:"url"`
	User            *string  `json:"user" mapstructure:"user"`
	Character       string   `json:"character" mapstructure:"character"`
	URI             string   `json:"uri" mapstructure:"uri"`
	SimpleCharacter string   `json:"simpleCharacter" mapstructure:"simple_character"`
	Avatar          Avatar   `json:"avatar" mapstructure:"avatar"`
	SharingURL      string   `json:"sharingURL" mapstructure:"sharing_url"`
	Type            string   `json:"type" mapstructure:"type"`
	ID              string   `json:"id" mapstructure:"id"`
	LatinName       string   `json:"latinName" mapstructure:"latin_name"`
}

type Actor struct {
	Name            string   `json:"name" mapstructure:"name"`
	Roles           []string `json:"roles" mapstructure:"roles"`
	Title           string   `json:"title" mapstructure:"title"`
	URL             string   `json:"url" mapstructure:"url"`
	User            *User    `json:"user" mapstructure:"user"`
	Character       string   `json:"character" mapstructure:"character"`
	URI             string   `json:"uri" mapstructure:"uri"`
	SimpleCharacter string   `json:"simpleCharacter" mapstructure:"simple_character"`
	Avatar          Avatar   `json:"avatar" mapstructure:"avatar"`
	SharingURL      string   `json:"sharingURL" mapstructure:"sharing_url"`
	Type            string   `json:"type" mapstructure:"type"`
	ID              string   `json:"id" mapstructure:"id"`
	LatinName       string   `json:"latinName" mapstructure:"latin_name"`
}
