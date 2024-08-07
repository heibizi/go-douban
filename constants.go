package douban

type InterestsType string

const (
	InterestsTypeWish    InterestsType = "wish"
	InterestsTypeCollect InterestsType = "collect"
	InterestsTypeDo      InterestsType = "do"
)

var Movie = MediaType{
	Code: "movie",
	Name: "电影",
}

var Tv = MediaType{
	Code: "tv",
	Name: "电视剧",
}

var MediaTypes = []MediaType{
	Movie,
	Tv,
}
