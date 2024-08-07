package douban

import (
	"github.com/heibizi/go-siteadapt"
	"strconv"
)

var (
	// SearchSubjectsTypeMovieNew 最新电影
	SearchSubjectsTypeMovieNew = SearchSubjectsType{
		MediaType: Movie,
		Tag:       "最新",
	}
	// SearchSubjectsTypeMovieHot 热门电影
	SearchSubjectsTypeMovieHot = SearchSubjectsType{
		MediaType: Movie,
		Tag:       "热门",
	}
	// SearchSubjectsTypeMovieRate 豆瓣高分
	SearchSubjectsTypeMovieRate = SearchSubjectsType{
		MediaType: Movie,
		Tag:       "豆瓣高分",
	}
	// SearchSubjectsTypeTvHot 热门电视剧
	SearchSubjectsTypeTvHot = SearchSubjectsType{
		MediaType: Tv,
		Tag:       "热门",
	}
	// SearchSubjectsTypeAnimeHot 热门动漫
	SearchSubjectsTypeAnimeHot = SearchSubjectsType{
		MediaType: Tv,
		Tag:       "日本动画",
	}
	// SearchSubjectsTypeVarietyHot 热门综艺
	SearchSubjectsTypeVarietyHot = SearchSubjectsType{
		MediaType: Tv,
		Tag:       "综艺",
	}
)

func (c *WebClient) SearchSubjects(mediaType MediaType, tag string, pageLimit int, pageStart int) ([]SearchSubject, error) {
	var d []SearchSubject
	err := siteadapt.NewSiteAdaptor(doubanWebConfig).List(siteadapt.RequestSiteParams{
		ReqId: "search_subjects",
		Env:   map[string]string{"type": mediaType.Code, "tag": tag, "page_limit": strconv.Itoa(pageLimit), "page_start": strconv.Itoa(pageStart)},
		UA:    c.ua,
	}, &d, nil)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (c *WebClient) SearchSubjectsWith(typ SearchSubjectsType, pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjects(typ.MediaType, typ.Tag, pageLimit, pageStart)
}

// MovieNew 最新电影
func (c *WebClient) MovieNew(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeMovieNew, pageLimit, pageStart)
}

// MovieHot 热门电影
func (c *WebClient) MovieHot(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeMovieHot, pageLimit, pageStart)
}

// MovieRate 高分电影
func (c *WebClient) MovieRate(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeMovieRate, pageLimit, pageStart)
}

// TvHot 热门电视剧
func (c *WebClient) TvHot(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeTvHot, pageLimit, pageStart)
}

// AnimeHot 热门动漫
func (c *WebClient) AnimeHot(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeAnimeHot, pageLimit, pageStart)
}

// VarietyHot 热门综艺
func (c *WebClient) VarietyHot(pageLimit int, pageStart int) ([]SearchSubject, error) {
	return c.SearchSubjectsWith(SearchSubjectsTypeVarietyHot, pageLimit, pageStart)
}
