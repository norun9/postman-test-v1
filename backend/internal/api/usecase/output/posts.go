package output

import (
	"github.com/norun9/postmantest/internal/api/domain/model"
)

type ChildPost struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func GetChildPost(post *model.Post) *ChildPost {
	return &ChildPost{
		ID:      post.ID,
		Content: post.Content.String,
	}
}
