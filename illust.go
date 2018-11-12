package pixiv

import (
	"context"
	"encoding/json"
	"io"
)

type ImageUrls map[string]string

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	// html(like?) document
	Caption   string    `json:"caption"`
	Restrict  int       `json:"restrict"`
	User      User      `json:"user"`
	Tags      []Tag     `json:"tags"`
	ImageUrls ImageUrls `json:"image_urls"`
	// use time.RFC3339 to decode
	CreateDate     string `json:"create_date"`
	PageCount      int    `json:"page_count"`
	TotalView      int    `json:"total_view"`
	TotalBookmarks int    `json:"total_bookmarks"`
	IsBookmarked   bool   `json:"is_bookmarked"`
	IsMuted        bool   `json:"is_muted"`
	Visible        bool   `json:"visible"`
}

type Illust struct {
	Item
	// illust / manga
	Type  string   `json:"type"`
	Tools []string `json:"tools"`

	Width  int `json:"width"`
	Height int `json:"height"`
	// Series <unknown type> `json:"series"`
	SanityLevel int `json:"sanity_level"`
	// appears when page_count > 1
	MetaPages []*MetaPage `json:"meta_pages"`
	// appears when page_count == 1
	MetaSinglePage *MetaSinglePage `json:"meta_single_page"`
}

type Novel struct {
	Item
	TextLength int `json:"text_length"`
}

type Tag struct {
	Name string `json:"name"`
}

type MetaPage struct {
	ImageUrls ImageUrls `json:"image_urls"`
}

type MetaSinglePage struct {
	OriginalImageUrl string `json:"original_image_url"`
}

func UnmarshalIllustList(r io.Reader) ([]*Illust, string, error) {
	decoder := json.NewDecoder(r)
	// TODO: full extract?
	d := struct {
		Illusts []*Illust `json:"illusts"`
		NextUrl string    `json:"next_url"`
	}{}
	if err := decoder.Decode(&d); err != nil {
		return nil, "", err
	}
	return d.Illusts, d.NextUrl, nil
}

func FetchIllustList(ctx context.Context, session Session, path, query string) ([]*Illust, string, error) {
	resp, err := session.Get(ctx, path, query)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	return UnmarshalIllustList(resp.Body)
}

func UnmarshalCommentList(r io.Reader) ([]*Comment, string, error) {
	decoder := json.NewDecoder(r)
	// TODO: full extract?
	d := struct {
		Illusts []*Comment `json:"comments"`
		NextUrl string     `json:"next_url"`
	}{}
	if err := decoder.Decode(&d); err != nil {
		return nil, "", err
	}
	return d.Illusts, d.NextUrl, nil
}

func FetchCommentList(ctx context.Context, session Session, path, query string) ([]*Comment, string, error) {
	resp, err := session.Get(ctx, path, query)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	return UnmarshalCommentList(resp.Body)
}

func FetchUserPreviewList(ctx context.Context, session Session, path, query string) ([]*UserPreview, string, error) {
	resp, err := session.Get(ctx, path, query)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	return UnmarshalUserPreviewList(resp.Body)
}

func UnmarshalUserPreviewList(r io.Reader) ([]*UserPreview, string, error) {
	decoder := json.NewDecoder(r)
	// TODO: full extract?
	d := struct {
		UserPreviews []*UserPreview `json:"user_previews"`
		NextUrl      string         `json:"next_url"`
	}{}
	if err := decoder.Decode(&d); err != nil {
		return nil, "", err
	}
	return d.UserPreviews, d.NextUrl, nil
}
