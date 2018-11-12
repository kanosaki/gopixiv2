package pixiv

type Comment struct {
	ID         uint64 `json:"id"`
	Comment    string `json:"comment"`
	Date       string `json:"date"`
	User       *User  `json:"user"`
	HasReplies bool   `json:"has_replies"`
}
