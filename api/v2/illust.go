package v2

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"go-pixiv"
)

type Illust struct {
	BasePath string
	client   pixiv.Session
}

func NewIllust(client pixiv.Session) *Illust {
	return &Illust{
		BasePath: "/v2/illust",
		client:   client,
	}
}

type IllustFollowQuery struct {
	// all, (private?, public?)
	Restrict string `url:"restrict"`
	Offset   int    `url:"offset,omitempty"`
}

func (i *Illust) Follow(ctx context.Context, q *IllustFollowQuery) ([]*pixiv.Illust, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchIllustList(ctx, i.client, i.BasePath+"/follow", v.Encode())
}

func (i *Illust) Related(ctx context.Context, illustID uint64) ([]*pixiv.Illust, string, error) {
	return pixiv.FetchIllustList(ctx, i.client, i.BasePath+"/related", fmt.Sprintf("illust_id=%d", illustID))
}

func (i *Illust) Comments(ctx context.Context, illustID uint64) ([]*pixiv.Comment, string, error) {
	return pixiv.FetchCommentList(ctx, i.client, i.BasePath+"/comments", fmt.Sprintf("illust_id=%d", illustID))
}
