package movie

type DoubanMovie struct {
	Title    string
	SubTitle string
	Alias    string
	Desc     string
	Year     int
	Area     string
	Tag      string
	Star     int
	Score    float64
	Quote    string
	DtCreate string
}

type Page struct {
	Page int
	Url string
}
