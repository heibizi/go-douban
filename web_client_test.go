package douban_test

import (
	"encoding/json"
	"github.com/heibizi/go-douban"
	"os"
	"testing"
)

var cookie = os.Getenv("GO_DOUBAN_COOKIE")
var ua = os.Getenv("GO_DOUBAN_UA")
var peopleId = os.Getenv("GO_DOUBAN_PEOPLE_ID")

const id = "26816104"

func TestMain(m *testing.M) {
	m.Run()
}

func TestDetail(t *testing.T) {
	d, err := newClient().Detail(id)
	log(d, err, t)
}

func TestInterestsRss(t *testing.T) {
	client := newClient()
	d, err := client.InterestsRss(peopleId)
	log(d, err, t)
}

func TestNowPlaying(t *testing.T) {
	d, err := newClient().NowPlaying()
	log(d, err, t)
}

func TestLater(t *testing.T) {
	d, err := newClient().Later()
	log(d, err, t)
}

func TestTop250(t *testing.T) {
	d, err := newClient().Top250()
	log(d, err, t)
}

func TestCollect(t *testing.T) {
	d, err := newClient().Collect(peopleId, 0)
	log(d, err, t)
}

func TestWish(t *testing.T) {
	d, err := newClient().Wish(peopleId, 0)
	log(d, err, t)
}

func TestDo(t *testing.T) {
	d, err := newClient().Do(peopleId, 0)
	log(d, err, t)
}

func TestSearchSubjects(t *testing.T) {
	d, err := newClient().SearchSubjects(douban.Movie, "最新", 10, 0)
	log(d, err, t)
}

func TestMovieNew(t *testing.T) {
	d, err := newClient().MovieNew(10, 0)
	log(d, err, t)
}

func TestMovieHot(t *testing.T) {
	d, err := newClient().MovieHot(10, 0)
	log(d, err, t)
}

func TestMovieRate(t *testing.T) {
	d, err := newClient().MovieRate(10, 0)
	log(d, err, t)
}

func TestTvHot(t *testing.T) {
	d, err := newClient().TvHot(10, 0)
	log(d, err, t)
}

func TestAnimeHot(t *testing.T) {
	d, err := newClient().AnimeHot(10, 0)
	log(d, err, t)
}

func TestVarietyHot(t *testing.T) {
	d, err := newClient().VarietyHot(10, 0)
	log(d, err, t)
}

func newClient() *douban.WebClient {
	client := douban.NewWebClient(cookie, ua)
	return client
}

func log(v any, err error, t *testing.T) {
	if err != nil {
		t.Log(err)
		return
	}
	j, _ := json.MarshalIndent(v, "", "    ")
	t.Log(string(j))
}
