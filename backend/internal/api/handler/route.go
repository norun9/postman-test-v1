package handler

import "github.com/norun9/postmantest/internal/api/usecase"

type Method string
type Permission string

var (
	Get Method = "GET"
)

type Path struct {
	Path string
	Method
}

type Route struct {
	Name string
	Func interface{}
}

func GetRouteMap(
	postUsecase usecase.Post,
) map[Path]Route {
	return map[Path]Route{
		{"/v1/posts/{id}", Get}: {"GET_POST_BY_ID", postUsecase.Get},
	}
}
