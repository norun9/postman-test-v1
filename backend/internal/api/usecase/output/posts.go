package output

import (
	"time"

	"github.com/norun9/postmantest/internal/api/domain/model"
	"github.com/volatiletech/null/v8"
)

type ChildPost struct {
	ID        int64       `json:"id"`
	Content   null.String `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func GetChildPost(post *model.Post) *ChildPost {
	return &ChildPost{
		ID:        post.ID,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
