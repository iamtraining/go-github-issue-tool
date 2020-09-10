package entity

import (
	"fmt"
	"time"
)

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int    `json:"number,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	Title     string `json:"title,omitempty"`
	User      *User
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	State     string    `json:"state,omitempty"`
}

func (i Issue) String() string {
	return fmt.Sprintf(`Issue %d (%s)
Title: %s State: %s
Login: %s Time: %v
Body: %s`, i.Number, i.HTMLURL, i.Title, i.State, i.User.Login, i.CreatedAt.Format(time.RFC822), i.Body)
}
