package query

import (
	"context"

	"github.com/norun9/postmantest/internal/api/domain/model"
)

type Post interface {
	GetByID(ctx context.Context, id int64) (reuslt *model.Post, err error)
}
