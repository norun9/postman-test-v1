package usecase

import (
	"context"

	"github.com/norun9/postmantest/internal/api/usecase/input"
	"github.com/norun9/postmantest/internal/api/usecase/output"
)

type Post interface {
	Get(ctx context.Context, p input.GetPost) (result *output.ChildPost, err error)
}
