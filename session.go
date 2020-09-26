package pixiv

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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

func NewOAuthClient(auth *AuthConfig) (*OAuthSession, error) {
	config := &oauth2.Config{
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "", // never used
			TokenURL: "https://oauth.secure.pixiv.net/auth/token",
		},
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &apiTransport{
			upstream:        http.DefaultTransport,
			signatureSecret: auth.SignatureSecret,
		},
	})
	ts := &OAuthTokenSource{
		config: config,
		auth:   auth,
		ctx:    ctx,
	}
	return &OAuthSession{
		client:   oauth2.NewClient(ctx, ts),
		Endpoint: DefaultAPIEndpoint,
		Limiter:  rate.NewLimiter(rate.Every(DefaultRateLimitEvery), DefaultRateLimitBurst),
	}, nil
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
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SignatureSecret string `json:"signature_secret"`
}

type OAuthTokenSource struct {
	config *oauth2.Config
	auth   *AuthConfig
	ctx    context.Context
}

func (o *OAuthTokenSource) Token() (*oauth2.Token, error) {
	return o.config.PasswordCredentialsToken(o.ctx, o.auth.Username, o.auth.Password)
}

type apiTransport struct {
	signatureSecret string
	upstream        http.RoundTripper
}

func (a *apiTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	timeString, hash := computeSignature(a.signatureSecret, time.Now())
	r.Header.Set("User-Agent", "PixivIOSApp/7.8.19 (iOS 13.3.1; iPhone11,2)")
	r.Header.Set("X-Client-Time", timeString)
	r.Header.Set("X-Client-Hash", hash)
	return a.upstream.RoundTrip(r)
}

func computeSignature(secret string, t time.Time) (string, string) {
	timeString := t.Format(time.RFC3339)
	hashContent := timeString + secret
	hash := md5.Sum([]byte(hashContent))
	return timeString, hex.EncodeToString(hash[:])
}
