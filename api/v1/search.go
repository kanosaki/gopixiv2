package v1

import (
	"context"
	"github.com/google/go-querystring/query"
	"go-pixiv"
)

type Search struct {
	BasePath string
	client   pixiv.Session
}

func NewSearch(client pixiv.Session) *Search {
	return &Search{
		BasePath: "/v1/search",
		client:   client,
	}
}

type SearchIllustQuery struct {
	Word         string       `url:"word"`
	SearchTarget SearchTarget `url:"search_target"`

	Sort           SearchSortOrder `url:"sort,omitempty"`
	BookmarkNumMin int             `url:"bookmark_num_min,omitempty"`
	BookmarkNumMax int             `url:"bookmark_num_max,omitempty"`
	// yyyy-mm-dd (2006-01-02)
	StartDate string `url:"start_date,omitempty"`
	// yyyy-mm-dd (2006-01-02)
	EndDate string `url:"end_date,omitempty"`
}

func (s *Search) Illust(ctx context.Context, q *SearchIllustQuery) ([]*pixiv.Illust, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchIllustList(ctx, s.client, s.BasePath+"/illust", v.Encode())
}
