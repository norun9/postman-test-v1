package query

import (
	"context"
	"database/sql"

	"github.com/norun9/postmantest/internal/api/domain/model"
	"github.com/norun9/postmantest/internal/api/domain/query"
	"github.com/norun9/postmantest/internal/api/infra/transfer"
	"github.com/norun9/postmantest/pkg/db"
	"github.com/norun9/postmantest/pkg/dbmodels"
	"github.com/norun9/postmantest/pkg/errof"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil/qm"
)

type post struct {
	dbClient db.Client
}

func NewPost(dbClient db.Client) query.Post {
	return &post{dbClient}
}

// GetByID :
func (q *post) GetByID(ctx context.Context, id int64) (reuslt *model.Post, err error) {
	queries := []qm.QueryMod{
		dbmodels.PostWhere.ID.EQ(id),
	}
	var dbPost *dbmodels.Post
	if dbPost, err = dbmodels.Posts(queries...).One(ctx, q.dbClient.Get(ctx)); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WithStack(errof.ErrNoPost)
		}
		return nil, errors.Wrap(errof.ErrDatabase, err.Error())
	}
	return transfer.ToPostDomain(dbPost), nil
}
