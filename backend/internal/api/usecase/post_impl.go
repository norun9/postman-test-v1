package usecase

import (
	"context"

	"github.com/norun9/postmantest/internal/api/domain/model"
	"github.com/norun9/postmantest/internal/api/domain/query"
	"github.com/norun9/postmantest/internal/api/usecase/input"
	"github.com/norun9/postmantest/internal/api/usecase/output"
)

type post struct {
	postQuery query.Post
}

func NewPost(
	postQuery query.Post,
) Post {
	return &post{
		postQuery: postQuery,
	}
}

func (u *post) Get(ctx context.Context, p input.GetPost) (result *output.ChildPost, err error) {
	var domainPost *model.Post
	if domainPost, err = u.postQuery.GetByID(ctx, p.ID); err != nil {
		return nil, err
	}
	return output.GetChildPost(domainPost), nil
}
