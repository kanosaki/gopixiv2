package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kanosaki/gopixiv2"
)

type Illust struct {
	BasePath string
	client   pixiv.Session
}

func NewIllust(client pixiv.Session) *Illust {
	return &Illust{
		BasePath: "/v1/illust",
		client:   client,
	}
}

type RankingQuery struct {
	Mode RankingMode `url:"mode"`
	// yyyy-mm-dd  (2006-01-02)
	Date   string `url:"date,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

func (i *Illust) Ranking(ctx context.Context, q *RankingQuery) ([]*pixiv.Illust, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchIllustList(ctx, i.client, i.BasePath+"/ranking", v.Encode())
}

type IllustRecommendedQuery struct {
	MinBookmarkIDForRecentIllust uint64 `url:"min_bookmark_id_for_recent_illust,omitempty"`
	MaxBookmarkIDForRecommend    uint64 `url:"max_bookmark_id_for_recommend,omitempty"`
	Offset                       int    `url:"offset,omitempty"`
}

func (i *Illust) Recommended(ctx context.Context, q *IllustRecommendedQuery) ([]*pixiv.Illust, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchIllustList(ctx, i.client, i.BasePath+"/recommended", v.Encode())
}

func (i *Illust) Detail(ctx context.Context, illustID uint64) (*pixiv.Illust, error) {
	resp, err := i.client.Get(ctx, i.BasePath+"/detail", fmt.Sprintf("illust_id=%d", illustID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	d := struct {
		Illust *pixiv.Illust `json:"illust"`
	}{}
	if err := decoder.Decode(&d); err != nil {
		return nil, err
	}
	return d.Illust, nil
}
