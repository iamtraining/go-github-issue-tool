package entity

import "time"

const (
	repoUser = "https://api.github.com/repos/%s/%s/issues"
	issue    = repoUser + "/%s"
)

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number  int    `json:"number,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
	Title   string `json:"title,omitempty"`
	*User
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	State     string    `json:"state,omitempty"`
}
