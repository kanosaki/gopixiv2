package pixiv

type User struct {
	ID               uint32    `json:"id"`
	Name             string    `json:"name"`
	Account          string    `json:"account"`
	ProfileImageUrls ImageUrls `json:"profile_image_urls"`

	IsFollowed bool   `json:"is_followed,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type UserPreview struct {
	User    *User     `json:"user"`
	Illusts []*Illust `json:"illusts"`
	Novels  []*Novel  `json:"novels"`
	IsMuted bool      `json:"is_muted"`
}

type UserDetails struct {
	User             *User                  `json:"user"`
	Profile          map[string]interface{} `json:"profile"`
	ProfilePublicity map[string]interface{} `json:"profile_publicity"`
	Workspace        map[string]interface{} `json:"workspace"`
}
