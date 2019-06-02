package pixiv

import (
	"github.com/kanosaki/pixiv_oauth2"
)

type OAuthClientOption func(*OAuthSession) error

func WithTokenCache(cachePath string) OAuthClientOption {
	return func(session *OAuthSession) error {
		transport, ok := session.client.Transport.(*pixiv_oauth2.Transport)
		if !ok {
			panic("never here")
		}
		origin := transport.Source
		transport.Source = &CachedTokenSource{
			cachePath: cachePath,
			origin:    origin,
		}
		return nil
	}
}
