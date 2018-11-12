package v2

import (
	"context"
	"fmt"
	"go-pixiv"
	"strings"
)

type IllustBookmark struct {
	BasePath string
	client   pixiv.Session
}

func NewIllustBookmark(client pixiv.Session) *Illust {
	return &Illust{
		BasePath: "/v2/illust/bookmark",
		client:   client,
	}
}

// restrict :: public, private
func (ib *IllustBookmark) Add(ctx context.Context, illustID uint64, restrict string) error {
	bodyStr := fmt.Sprintf("illust_id=%d&restrict=%s", illustID, restrict)
	// json?, simulating app behavior
	resp, err := ib.client.Post(ctx, ib.BasePath+"/add", "", "application/json; charset=utf-8", strings.NewReader(bodyStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
