package v1

import (
	"context"
	"fmt"
	"github.com/kanosaki/gopixiv2"
	"github.com/kanosaki/gopixiv2/common"
	"strconv"
)

type API struct {
	Illust       *Illust
	User         *User
	Search       *Search
	TrendingTags *TrendingTags
}

func NewAPI(client pixiv.Session) *API {
	return &API{
		Illust:       NewIllust(client),
		User:         NewUser(client),
		Search:       NewSearch(client),
		TrendingTags: NewTrendingTags(client),
	}
}

func (a *API) DoGet(ctx context.Context, path string, params map[string]string) (interface{}, error) {
	dispatchers := map[string]func() (interface{}, error){
		"/v1/illust/ranking": func() (interface{}, error) {
			return common.WrapContinuous(a.Illust.Ranking(ctx, &RankingQuery{Mode: RankingMode(params["mode"])}))
		},
		"/v1/illust/recommended": func() (interface{}, error) {
			return common.WrapContinuous(a.Illust.Recommended(ctx, &IllustRecommendedQuery{}))
		},
		"/v1/illust/detail": func() (interface{}, error) {
			id, err := strconv.ParseUint(params["id"], 10, 64)
			if err != nil {
				return nil, err
			}
			return a.Illust.Detail(ctx, id)
		},
		"/v1/search/illust": func() (interface{}, error) {
			return common.WrapContinuous(a.Search.Illust(ctx, &SearchIllustQuery{
				Word:         params["word"],
				SearchTarget: SearchTargetPartialMatchForTags,
			}))
		},
		"/v1/user/illusts": func() (interface{}, error) {
			id, err := strconv.ParseUint(params["id"], 10, 64)
			if err != nil {
				return nil, err
			}
			return common.WrapContinuous(a.User.Illusts(ctx, &UserIllustQuery{UserID: id}))
		},
	}
	if d, ok := dispatchers[path]; ok {
		return d()
	} else {
		return nil, fmt.Errorf("not found (at /v1): %s", path)
	}
}
