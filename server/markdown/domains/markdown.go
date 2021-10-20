package markdownDomains

type Markdown struct {
	ID      string `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	UserID  string `json:"userId" db:"user_id"`
}
