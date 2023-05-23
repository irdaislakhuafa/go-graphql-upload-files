package entity

import "github.com/99designs/gqlgen/graphql"

type Todo struct {
	ID       string         `json:"id,omitempty"`
	Text     string         `json:"text,omitempty"`
	Done     bool           `json:"done,omitempty"`
	UserID   string         `json:"userID,omitempty"`
	File     graphql.Upload `json:"file,omitempty"`
	FilePath string         `json:"filePath,omitempty"`
}
