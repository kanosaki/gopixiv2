package pixiv

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	oauth2 "github.com/kanosaki/pixiv_oauth2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const (
	DefaultAPIEndpoint    = "https://app-api.pixiv.net"
	DefaultRateLimitEvery = 1 * time.Second
	DefaultRateLimitBurst = 10
)

type Session interface {
	Get(ctx context.Context, base, query string) (*http.Response, error)
	Post(ctx context.Context, base, query, contentType string, body io.Reader) (*http.Response, error)
}

// A client for Pixiv smartphone apis
type OAuthSession struct {
	client   *http.Client
	Endpoint string
	Limiter  *rate.Limiter
}

func NewOAuthClient(auth *AuthConfig, opts ...OAuthClientOption) (*OAuthSession, error) {
	config := &oauth2.Config{
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "", // never used
			TokenURL: "https://oauth.secure.pixiv.net/auth/token",
		},
	}
	ctx := context.Background()
	ts := &OAuthTokenSource{
		config: config,
		auth:   auth,
		ctx:    ctx,
	}
	sess := &OAuthSession{
		client:   oauth2.NewClient(ctx, ts),
		Endpoint: DefaultAPIEndpoint,
		Limiter:  rate.NewLimiter(rate.Every(DefaultRateLimitEvery), DefaultRateLimitBurst),
	}
	for _, opt := range opts {
		if err := opt(sess); err != nil {
			return nil, err
		}
	}
	return sess, nil
}

func (s *OAuthSession) Post(ctx context.Context, base, query, contentType string, body io.Reader) (*http.Response, error) {
	if err := s.Limiter.Wait(ctx); err != nil {
		return nil, err
	}
	u := fmt.Sprintf("%s/%s", s.Endpoint, base)
	if query != "" {
		u += "?" + query
	}
	log.Debug("POST ", u)
	return s.client.Post(u, contentType, body)
}

func (s *OAuthSession) Get(ctx context.Context, base, query string) (*http.Response, error) {
	if err := s.Limiter.Wait(ctx); err != nil {
		return nil, err
	}
	u := s.Endpoint + base
	if query != "" {
		u += "?" + query
	}
	log.Debug("GET ", u)
	return s.client.Get(u)
}

type AuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type OAuthTokenSource struct {
	config *oauth2.Config
	auth   *AuthConfig
	ctx    context.Context
}

func (o *OAuthTokenSource) Token() (*oauth2.Token, error) {
	return o.config.PasswordCredentialsToken(o.ctx, o.auth.Username, o.auth.Password)
}
