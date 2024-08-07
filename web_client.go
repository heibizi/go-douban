package douban

import (
	"encoding/xml"
	"fmt"
	"github.com/heibizi/go-siteadapt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func loadConfig(file string) siteadapt.Config {
	_, filename, _, _ := runtime.Caller(0)
	f, _ := os.Open(filepath.Join(filepath.Dir(filename), file))
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	b, _ := io.ReadAll(f)
	c, err := siteadapt.NewConfigReader(b).Read()
	if err != nil {
		panic(err)
	}
	return *c
}

var doubanWebConfig = loadConfig("web_config.json")

type WebClient struct {
	cookie string
	ua     string
}

func NewWebClient(cookie string, ua string) *WebClient {
	return &WebClient{cookie, ua}
}

// Detail 获取豆瓣详情
func (c *WebClient) Detail(id string) (Detail, error) {
	var d Detail
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).Data(siteadapt.RequestSiteParams{
		ReqId:  "detail",
		Env:    map[string]string{"id": id},
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// InterestsRss 获取用户动态
func (c *WebClient) InterestsRss(peopleId string) ([]InterestsRssInfo, error) {
	var data []byte
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).Raw(siteadapt.RequestSiteParams{
		ReqId:  "interests_rss",
		Env:    map[string]string{"peopleId": peopleId},
		Cookie: c.cookie,
		UA:     c.ua,
	}, func(result siteadapt.RawResult) {
		data = result.Data
	})
	if err != nil {
		return nil, err
	}
	var rss interestsRssResult
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, fmt.Errorf("解析用户动态异常")
	}
	var d []InterestsRssInfo
	for _, item := range rss.Items {
		rssType := ""
		if len(item.Title) > 2 {
			rssType = string([]rune(item.Title)[0:2])
		}
		title := string([]rune(item.Title)[2:])
		info := InterestsRssInfo{
			Title: title,
			Url:   item.Link,
			Date:  item.PubDate,
		}
		if rssType == "想看" {
			info.Type = InterestsTypeWish
		} else if rssType == "看过" {
			info.Type = InterestsTypeCollect
		} else if rssType == "在看" {
			info.Type = InterestsTypeDo
		}
		date, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return nil, fmt.Errorf("解析日期异常: %v", err)
		}
		info.Date = date.Format("2006-01-02")
		d = append(d, info)
	}
	return d, nil
}

// NowPlaying 正在热映
func (c *WebClient) NowPlaying() ([]NowPlaying, error) {
	var d []NowPlaying
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "now_playing",
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Later 即将上映
func (c *WebClient) Later() ([]Later, error) {
	var d []Later
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "later",
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Top250 Top 250
func (c *WebClient) Top250() ([]Top250, error) {
	var d []Top250
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "top250",
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Collect 获取收藏
func (c *WebClient) Collect(peopleId string, start int) ([]Collect, error) {
	var d []Collect
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "collect",
		Env:    map[string]string{"peopleId": peopleId, "start": strconv.Itoa(start)},
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Wish 获取想看
func (c *WebClient) Wish(peopleId string, start int) ([]Wish, error) {
	var d []Wish
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "wish",
		Env:    map[string]string{"peopleId": peopleId, "start": strconv.Itoa(start)},
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Do 获取在看
func (c *WebClient) Do(peopleId string, start int) ([]Do, error) {
	var d []Do
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId:  "do",
		Env:    map[string]string{"peopleId": peopleId, "start": strconv.Itoa(start)},
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

// User 获取用户信息
func (c *WebClient) User(peopleId string) (UserInfo, error) {
	var d UserInfo
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).Data(siteadapt.RequestSiteParams{
		ReqId:  "user",
		Env:    map[string]string{"peopleId": peopleId},
		Cookie: c.cookie,
		UA:     c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}
