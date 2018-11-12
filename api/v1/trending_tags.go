package v1

import (
	"context"
	"encoding/json"
	"github.com/kanosaki/gopixiv2"
)

type TrendingTags struct {
	BasePath string
	client   pixiv.Session
}

func NewTrendingTags(client pixiv.Session) *TrendingTags {
	return &TrendingTags{
		BasePath: "/v1/trending-tags",
		client:   client,
	}
}

func (i *Illust) Illust(ctx context.Context) ([]*pixiv.TrendingTag, error) {
	resp, err := i.client.Get(ctx, i.BasePath+"/illust", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	d := struct {
		TrendTags []*pixiv.TrendingTag `json:"trend_tags"`
	}{}
	if err := decoder.Decode(&d); err != nil {
		return nil, err
	}
	return d.TrendTags, nil
}
