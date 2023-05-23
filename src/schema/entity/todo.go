package entity

import "github.com/99designs/gqlgen/graphql"

type Todo struct {
	ID     string `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Done   bool   `json:"done,omitempty"`
	UserID string `json:"userID,omitempty"`
	FileID string `json:"fileID,omitempty"`
	File   graphql.Upload
}
