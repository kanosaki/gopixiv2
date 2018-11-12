package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"go-pixiv"
)

type User struct {
	BasePath string
	client   pixiv.Session
}

func NewUser(client pixiv.Session) *User {
	return &User{
		BasePath: "/v1/user",
		client:   client,
	}
}

type UserIllustQuery struct {
	UserID uint64 `url:"user_id"`
	// illust, manga, novel
	Type string `url:"type"`
}

func (u *User) Illusts(ctx context.Context, q *UserIllustQuery) ([]*pixiv.Illust, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchIllustList(ctx, u.client, u.BasePath+"/illusts", v.Encode())
}

type UserFollowingQuery struct {
	UserID uint64 `url:"user_id"`
	// public, private
	Restrict string `url:"restrict,omitempty"`
}

func (u *User) Following(ctx context.Context, q *UserFollowingQuery) ([]*pixiv.UserPreview, string, error) {
	v, _ := query.Values(q)
	return pixiv.FetchUserPreviewList(ctx, u.client, u.BasePath+"/following", v.Encode())
}

func (u *User) Detail(ctx context.Context, userID uint64) (*pixiv.UserDetails, error) {
	resp, err := u.client.Get(ctx, u.BasePath+"/detail", fmt.Sprintf("user_id=%d", userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	ud := &pixiv.UserDetails{}
	if err := decoder.Decode(&ud); err != nil {
		return nil, err
	}
	return ud, nil
}
