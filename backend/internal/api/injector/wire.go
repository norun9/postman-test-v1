// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/norun9/postmantest/internal/api/handler"
	"github.com/norun9/postmantest/internal/api/infra/query"
	"github.com/norun9/postmantest/internal/api/usecase"
	"github.com/norun9/postmantest/pkg/config"
	"github.com/norun9/postmantest/pkg/db"
)

var RouteMap = wire.NewSet(
	db.NewDB,
	db.NewMySQL,

	usecase.NewPost,

	query.NewPost,
)

func InitializeRestHandler(config.HTTP, config.MySQL) (_ handler.RestHandler) {
	wire.Build(
		handler.NewRestHandler,
		handler.GetRouteMap,
		RouteMap,
	)
	return
}

func InitializeRouteMap(config.HTTP, config.MySQL) (_ map[handler.Path]handler.Route) {
	wire.Build(
		handler.GetRouteMap,
		RouteMap,
	)
	return
}
