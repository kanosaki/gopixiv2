package pixiv

type TrendingTag struct {
	Tag    string  `json:"tag"`
	Illust *Illust `json:"illust"`
}
