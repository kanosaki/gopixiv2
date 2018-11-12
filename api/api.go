package api

import (
	"context"
	"fmt"
	"gopixiv2"
	"gopixiv2/api/v1"
	"gopixiv2/api/v2"
	"strings"
)

type API struct {
	V1 *v1.API
	V2 *v2.API
}

func New(client pixiv.Session) *API {
	return &API{
		V1: v1.NewAPI(client),
		V2: v2.NewAPI(client),
	}
}

func (a *API) DoGet(ctx context.Context, path string, params map[string]string) (interface{}, error) {
	if strings.HasPrefix(path, "/v1") {
		return a.V1.DoGet(ctx, path, params)
	} else if strings.HasPrefix(path, "/v2") {
		return a.V2.DoGet(ctx, path, params)
	} else {
		return nil, fmt.Errorf("not found(at /): %s", path)
	}
}
