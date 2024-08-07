package douban_test

import (
	"github.com/heibizi/go-douban"
	"testing"
)

const locId = "118172"
const douListId = "4026601"
const subjectCollectionId = "film_genre_44"

func newApiClient() *douban.ApiClient {
	return douban.NewApiClient()
}

func TestRankList(t *testing.T) {
	d, err := newApiClient().RankList(douban.Movie, 0, 20)
	log(d, err, t)
}

func TestYearRanks(t *testing.T) {
	d, err := newApiClient().YearRanks(douban.Movie, 2023)
	log(d, err, t)
}

func TestCategoryRanks(t *testing.T) {
	d, err := newApiClient().CategoryRanks(douban.Movie, 0, 20)
	log(d, err, t)
}

func TestMovieModules(t *testing.T) {
	d, err := newApiClient().Modules(douban.Movie)
	log(d, err, t)
}

func TestTvModules(t *testing.T) {
	d, err := newApiClient().Modules(douban.Tv)
	log(d, err, t)
}

func TestSearchWx(t *testing.T) {
	d, err := newApiClient().SearchWx("花千骨", 0, 20)
	log(d, err, t)
}

func TestSearch(t *testing.T) {
	d, err := newApiClient().Search("敦刻尔克", 0, 20)
	log(d, err, t)
}

func TestMovieRecommend(t *testing.T) {
	d, err := newApiClient().Recommend(douban.Movie, []string{"惊悚"}, "R", 0, 20)
	log(d, err, t)
}

func TestTvRecommend(t *testing.T) {
	d, err := newApiClient().Recommend(douban.Tv, []string{"惊悚"}, "R", 0, 20)
	log(d, err, t)
}

func TestTag(t *testing.T) {
	d, err := newApiClient().Tag(douban.Tv)
	log(d, err, t)
}

func TestSkyNetNewPlayLists(t *testing.T) {
	d, err := newApiClient().SkyNetNewPlayLists("all", douban.Tv, 0, 20)
	log(d, err, t)
}

func TestMovieShowing(t *testing.T) {
	d, err := newApiClient().MovieMovieShowing(locId, 0, 20, "recommend")
	log(d, err, t)
}

func TestHotCities(t *testing.T) {
	d, err := newApiClient().HotCities()
	log(d, err, t)
}

func TestMovieHotGaia(t *testing.T) {
	d, err := newApiClient().MovieHotGaia(locId, 0, 20, "recommend")
	log(d, err, t)
}

func TestComingSoon(t *testing.T) {
	d, err := newApiClient().ComingSoon(douban.Movie, 0, 20, "recommend", "华语")
	log(d, err, t)
}

func TestApiDetail(t *testing.T) {
	d, err := newApiClient().Detail(douban.Movie, id)
	log(d, err, t)
}

func TestRating(t *testing.T) {
	d, err := newApiClient().Rating(douban.Movie, id)
	log(d, err, t)
}

func TestPhotos(t *testing.T) {
	d, err := newApiClient().Photos(douban.Movie, id)
	log(d, err, t)
}

func TestTrailers(t *testing.T) {
	d, err := newApiClient().Trailers(douban.Movie, id)
	log(d, err, t)
}

func TestInterests(t *testing.T) {
	d, err := newApiClient().Interests(douban.Movie, id)
	log(d, err, t)
}

func TestReviews(t *testing.T) {
	d, err := newApiClient().Reviews(douban.Movie, id)
	log(d, err, t)
}

func TestRecommendations(t *testing.T) {
	d, err := newApiClient().Recommendations(douban.Movie, id)
	log(d, err, t)
}

func TestCelebrities(t *testing.T) {
	d, err := newApiClient().Celebrities(douban.Movie, id)
	log(d, err, t)
}

func TestDouList(t *testing.T) {
	d, err := newApiClient().DouList(douListId)
	log(d, err, t)
}

func TestDouListItems(t *testing.T) {
	d, err := newApiClient().DouListItems(douListId, 0, 20)
	log(d, err, t)
}

func TestSubjectCollection(t *testing.T) {
	d, err := newApiClient().SubjectCollection(subjectCollectionId)
	log(d, err, t)
}

func TestSubjectCollectionItems(t *testing.T) {
	d, err := newApiClient().SubjectCollectionItems(subjectCollectionId, 0, 20)
	log(d, err, t)
}

func TestSubjectCollectionItemsWith(t *testing.T) {
	d, err := newApiClient().SubjectCollectionItemsWith(douban.TvHot, 0, 20)
	log(d, err, t)
}
