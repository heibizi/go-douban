package douban

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"net/url"
)

func (c *ApiClient) Modules(mediaType MediaType) (Modules, error) {
	requestUrl, _ := url.JoinPath(apiUrl, "/"+mediaType.Code+"/modules")
	data, err := c.request(requestUrl, "GET", nil)
	var o Modules
	if err != nil {
		return o, fmt.Errorf("获取 modules 失败: %v", err)
	}
	result := gjson.Parse(string(data))
	var boards []SubjectCollectionBoards
	var collections []SelectedCollections
	var chartCollections []SelectedChartCollections
	var subjectEntrances SubjectEntrances
	var subjectUnions SubjectUnions
	var comingSoon ComingSoon
	var subjectSuggestion SubjectSuggestion
	comingSoonKey := ""
	if mediaType == Movie {
		comingSoonKey = "movie_coming_soon"
	} else if mediaType == Tv {
		comingSoonKey = "tv_coming_soon"
	}
	for _, module := range result.Get("modules").Array() {
		if module.Get("key").String() == "subject_entrances" {
			err = mapstructure.WeakDecode(module.Value(), &subjectEntrances)
			if err != nil {
				return o, fmt.Errorf("解析 subject_entrances 失败: %v", err)
			}
		} else if module.Get("key").String() == "subject_unions" {
			err = mapstructure.WeakDecode(module.Value(), &subjectUnions)
			if err != nil {
				return o, fmt.Errorf("解析 subject_unions 失败: %v", err)
			}
		} else if module.Get("key").String() == comingSoonKey {
			err = mapstructure.WeakDecode(module.Value(), &comingSoon)
			if err != nil {
				return o, fmt.Errorf("解析 coming_soon 失败: %v", err)
			}
		} else if module.Get("key").String() == "subject_collection_boards" {
			var subjectCollectionBoards SubjectCollectionBoards
			err = mapstructure.WeakDecode(module.Value(), &subjectCollectionBoards)
			if err != nil {
				return o, fmt.Errorf("解析 subject_suggestion 失败: %v", err)
			}
			boards = append(boards, subjectCollectionBoards)
		} else if module.Get("key").String() == "selected_collections" {
			var selectedCollections SelectedCollections
			err = mapstructure.WeakDecode(module.Value(), &selectedCollections)
			if err != nil {
				return o, fmt.Errorf("解析 selected_collections 失败: %v", err)
			}
			collections = append(collections, selectedCollections)
		} else if module.Get("key").String() == "selected_chart_collections" {
			var selectedChartCollections SelectedChartCollections
			err = mapstructure.WeakDecode(module.Value(), &selectedChartCollections)
			if err != nil {
				return o, fmt.Errorf("解析 selected_chart_collections 失败: %v", err)
			}
			chartCollections = append(chartCollections, selectedChartCollections)
		} else if module.Get("key").String() == "subject_suggestion" {
			err = mapstructure.WeakDecode(module.Value(), &subjectSuggestion)
			if err != nil {
				return o, fmt.Errorf("解析 subject_suggestion 失败: %v", err)
			}
		}
	}
	o.SubjectEntrances = subjectEntrances
	o.SubjectUnions = subjectUnions
	o.ComingSoon = comingSoon
	o.SubjectCollectionBoards = boards
	o.SelectedCollections = collections
	o.SelectedChartCollections = chartCollections
	o.SubjectSuggestion = subjectSuggestion
	return o, nil
}

type Modules struct {
	SubjectEntrances         SubjectEntrances           `json:"subjectEntrances" mapstructure:"subject_entrances"`
	SubjectUnions            SubjectUnions              `json:"subjectUnions" mapstructure:"subject_unions"`
	ComingSoon               ComingSoon                 `json:"comingSoon" mapstructure:"coming_soon"`
	SubjectCollectionBoards  []SubjectCollectionBoards  `json:"subjectCollectionBoards" mapstructure:"subject_collection_boards"`
	SelectedCollections      []SelectedCollections      `json:"selectedCollections" mapstructure:"selected_collections"`
	SelectedChartCollections []SelectedChartCollections `json:"selectedChartCollections" mapstructure:"selected_chart_collections"`
	SubjectSuggestion        SubjectSuggestion          `json:"subjectSuggestion" mapstructure:"subject_suggestion"`
}

type MovieModules struct {
	SubjectEntrances         SubjectEntrances         `json:"subjectEntrances" mapstructure:"subject_entrances"`
	SubjectUnions            SubjectUnions            `json:"subjectUnions" mapstructure:"subject_unions"`
	ComingSoon               ComingSoon               `json:"comingSoon" mapstructure:"coming_soon"`
	HotGaia                  SubjectCollectionBoards  `json:"hotGaia" mapstructure:"hot_gaia"`
	SelectedCollections      SelectedCollections      `json:"selectedCollections" mapstructure:"selected_collections"`
	SelectedChartCollections SelectedChartCollections `json:"selectedChartCollections" mapstructure:"selected_chart_collections"`
	MovieShowing             SubjectCollectionBoards  `json:"movieShowing" mapstructure:"movie_showing"`
	SubjectSuggestion        SubjectSuggestion        `json:"subjectSuggestion" mapstructure:"subject_suggestion"`
}

type (
	SubjectEntrances struct {
		URL        string      `json:"url" mapstructure:"url"`
		ModuleName string      `json:"moduleName" mapstructure:"module_name"`
		Data       SubjectData `json:"data" mapstructure:"data"`
		URI        string      `json:"uri" mapstructure:"uri"`
		Key        string      `json:"key" mapstructure:"key"`
	}

	SubjectData struct {
		SubjectEntrances []SubjectEntrance `json:"subjectEntrances" mapstructure:"subject_entrances"`
		Total            int               `json:"total" mapstructure:"total"`
	}

	SubjectEntrance struct {
		Icon  Icon   `json:"icon" mapstructure:"icon"`
		URI   string `json:"uri" mapstructure:"uri"`
		Key   string `json:"key" mapstructure:"key"`
		Title string `json:"title" mapstructure:"title"`
	}

	Icon struct {
		Dark  string `json:"dark" mapstructure:"dark"`
		Light string `json:"light" mapstructure:"light"`
	}

	SubjectUnions struct {
		ModuleName string      `json:"moduleName" mapstructure:"module_name"`
		Data       []UnionData `json:"data" mapstructure:"data"`
		Key        string      `json:"key" mapstructure:"key"`
	}

	UnionData struct {
		Type        string   `json:"type" mapstructure:"type"`
		Extra       []string `json:"extra" mapstructure:"extra"`
		Collections []string `json:"collections" mapstructure:"collections"`
		Title       string   `json:"title" mapstructure:"title"`
	}

	SubjectAd struct {
		URL        string `json:"url" mapstructure:"url"`
		ModuleName string `json:"moduleName" mapstructure:"module_name"`
		Data       AdData `json:"data" mapstructure:"data"`
		URI        string `json:"uri" mapstructure:"uri"`
		Key        string `json:"key" mapstructure:"key"`
	}

	AdData struct {
		Type string `json:"type" mapstructure:"type"`
		Unit string `json:"unit" mapstructure:"unit"`
	}

	ComingSoon struct {
		URL        string           `json:"url" mapstructure:"url"`
		ModuleName string           `json:"moduleName" mapstructure:"module_name"`
		Data       []ComingSoonData `json:"data" mapstructure:"data"`
		URI        string           `json:"uri" mapstructure:"uri"`
		Key        string           `json:"key" mapstructure:"key"`
	}

	ComingSoonData struct {
		Title    string     `json:"title" mapstructure:"title"`
		URI      string     `json:"uri" mapstructure:"uri"`
		TotalHot int        `json:"total_hot" mapstructure:"total_hot"`
		Total    int        `json:"total" mapstructure:"total"`
		Type     string     `json:"type" mapstructure:"type"`
		Items    []RankItem `json:"items" mapstructure:"items"`
	}

	SubjectCollectionBoards struct {
		URL        string `json:"url" mapstructure:"url"`
		ModuleName string `json:"moduleName" mapstructure:"module_name"`
		Data       struct {
			SubjectCollectionBoards []SubjectCollectionBoard `json:"subjectCollectionBoards" mapstructure:"subject_collection_boards"`
		} `json:"data" mapstructure:"data"`
		URI string `json:"uri" mapstructure:"uri"`
		Key string `json:"key" mapstructure:"key"`
	}

	SubjectCollectionBoard struct {
		SubjectCollection SubjectCollection            `json:"subject_collection" mapstructure:"subject_collection"`
		Type              string                       `json:"type" mapstructure:"type"`
		Items             []SubjectCollectionBoardItem `json:"items" mapstructure:"items"`
	}

	SubjectCollection struct {
		SubjectType           string                `json:"subjectType" mapstructure:"subject_type"`
		Subtitle              string                `json:"subtitle" mapstructure:"subtitle"`
		BackgroundColorScheme BackgroundColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
		SharingTitle          string                `json:"sharingTitle" mapstructure:"sharing_title"`
		UpdatedAt             *string               `json:"updatedAt" mapstructure:"updated_at"`
		ID                    string                `json:"id" mapstructure:"id"`
		ShowHeaderMask        bool                  `json:"showHeaderMask" mapstructure:"show_header_mask"`
		MediumName            string                `json:"mediumName" mapstructure:"medium_name"`
		Description           string                `json:"description" mapstructure:"description"`
		ShortName             string                `json:"shortName" mapstructure:"short_name"`
		NFollowers            *int                  `json:"nFollowers" mapstructure:"n_followers"`
		CoverURL              string                `json:"coverUrl" mapstructure:"cover_url"`
		ShowRank              bool                  `json:"showRank" mapstructure:"show_rank"`
		SharingURL            string                `json:"sharingUrl" mapstructure:"sharing_url"`
		SubjectCount          int                   `json:"subjectCount" mapstructure:"subject_count"`
		Name                  string                `json:"name" mapstructure:"name"`
		URL                   string                `json:"url" mapstructure:"url"`
		URI                   string                `json:"uri" mapstructure:"uri"`
		ShowFilterPlayable    bool                  `json:"showFilterPlayable" mapstructure:"show_filter_playable"`
		IconFgImage           string                `json:"iconFgImage" mapstructure:"icon_fg_image"`
		MoreDescription       string                `json:"moreDescription" mapstructure:"more_description"`
		Display               Display               `json:"display" mapstructure:"display"`
	}

	Cover struct {
		URL    string `json:"url" mapstructure:"url"`
		Width  int    `json:"width" mapstructure:"width"`
		Shape  string `json:"shape" mapstructure:"shape"`
		Height int    `json:"height" mapstructure:"height"`
	}

	SubjectCollectionBoardItem struct {
		OriginalPrice    *string  `json:"originalPrice" mapstructure:"original_price"`
		Rating           Rating   `json:"rating" mapstructure:"rating"`
		Actions          []string `json:"actions" mapstructure:"actions"`
		Year             string   `json:"year" mapstructure:"year"`
		CardSubtitle     string   `json:"cardSubtitle" mapstructure:"card_subtitle"`
		ID               string   `json:"id" mapstructure:"id"`
		Title            string   `json:"title" mapstructure:"title"`
		Label            *string  `json:"label" mapstructure:"label"`
		Actors           []string `json:"actors" mapstructure:"actors"`
		Interest         *string  `json:"interest" mapstructure:"interest"`
		Type             string   `json:"type" mapstructure:"type"`
		Description      string   `json:"description" mapstructure:"description"`
		HasLineWatch     bool     `json:"hasLineWatch" mapstructure:"has_linewatch"`
		Price            *string  `json:"price" mapstructure:"price"`
		Date             *string  `json:"date" mapstructure:"date"`
		Info             string   `json:"info" mapstructure:"info"`
		URL              string   `json:"url" mapstructure:"url"`
		ReleaseDate      string   `json:"releaseDate" mapstructure:"release_date"`
		Cover            Cover    `json:"cover" mapstructure:"cover"`
		URI              string   `json:"uri" mapstructure:"uri"`
		Subtype          string   `json:"subtype" mapstructure:"subtype"`
		Directors        []string `json:"directors" mapstructure:"directors"`
		ReviewerName     string   `json:"reviewerName" mapstructure:"reviewer_name"`
		NullRatingReason string   `json:"nullRatingReason" mapstructure:"null_rating_reason"`
	}

	SelectedChartCollections struct {
		URL        string `json:"url" mapstructure:"url"`
		ModuleName string `json:"moduleName" mapstructure:"module_name"`
		Data       struct {
			Total               int                       `json:"total" mapstructure:"total"`
			SelectedCollections []SelectedChartCollection `json:"selectedCollections" mapstructure:"selected_collections"`
			Title               string                    `json:"title" mapstructure:"title"`
		} `json:"data" mapstructure:"data"`
		URI string `json:"uri" mapstructure:"uri"`
		Key string `json:"key" mapstructure:"key"`
	}

	SelectedChartCollection struct {
		Category              string                `json:"category" mapstructure:"category"`
		ItemsCount            int                   `json:"itemsCount" mapstructure:"items_count"`
		IsMergedCover         bool                  `json:"isMergedCover" mapstructure:"is_merged_cover"`
		SharingURL            string                `json:"sharingUrl" mapstructure:"sharing_url"`
		Title                 string                `json:"title" mapstructure:"title"`
		URL                   string                `json:"url" mapstructure:"url"`
		URI                   string                `json:"uri" mapstructure:"uri"`
		CoverURL              string                `json:"coverUrl" mapstructure:"cover_url"`
		FollowersCount        int                   `json:"followersCount" mapstructure:"followers_count"`
		IsOfficial            bool                  `json:"isOfficial" mapstructure:"is_official"`
		Type                  string                `json:"type" mapstructure:"type"`
		ID                    string                `json:"id" mapstructure:"id"`
		DoneCount             int                   `json:"doneCount" mapstructure:"done_count"`
		BackgroundColorScheme BackgroundColorScheme `json:"backgroundColorScheme" mapstructure:"background_color_scheme"`
	}

	SubjectSuggestion struct {
		URL        string `json:"url" mapstructure:"url"`
		ModuleName string `json:"moduleName" mapstructure:"module_name"`
		Data       struct {
			Type  string `json:"type" mapstructure:"type"`
			Title string `json:"title" mapstructure:"title"`
		} `json:"data" mapstructure:"data"`
		URI string `json:"uri" mapstructure:"uri"`
		Key string `json:"key" mapstructure:"key"`
	}

	SelectedCollections struct {
		URL        string `json:"url" mapstructure:"url"`
		ModuleName string `json:"moduleName" mapstructure:"module_name"`
		Data       struct {
			Total               int                  `json:"total" mapstructure:"total"`
			SelectedCollections []SelectedCollection `json:"selectedCollections" mapstructure:"selected_collections"`
			Title               string               `json:"title" mapstructure:"title"`
		} `json:"data" mapstructure:"data"`
		URI string `json:"uri" mapstructure:"uri"`
		Key string `json:"key" mapstructure:"key"`
	}
)
