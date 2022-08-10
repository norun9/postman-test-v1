package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/inconshreveable/log15"
	"github.com/norun9/postmantest/pkg/errof"
	"github.com/pkg/errors"
)

var (
	idRegexp = regexp.MustCompile(`/[0-9]+(/|$)`)
)

// RestHandler :
type RestHandler interface {
	Exec(ctx context.Context, w http.ResponseWriter, r *http.Request, params interface{})
	GetRoute(router chi.Router)
	GetHealthRouter(router chi.Router)
}

type restHandler struct {
	routeMap map[Path]Route
}

// NewRestHandler :
func NewRestHandler(
	routeMap map[Path]Route,
) RestHandler {
	return restHandler{
		routeMap,
	}
}

func (h restHandler) GetRoute(r chi.Router) {
	// r.Use(middleware.Logger)
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()
	r.Use(middleware.Recoverer)
	r.Route("/v1/posts", func(r chi.Router) {
		r.Route("/", h.GetPostRouter)
	})
}

func (h restHandler) GetHealthRouter(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("status OK")); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
	})
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (h restHandler) Exec(ctx context.Context, w http.ResponseWriter, r *http.Request, params interface{}) {
	var err error
	var method Method
	defer func() {
		if p := recover(); p != nil {
			err = errors.Wrap(errof.ErrInternal, errof.PanicToErr(p).Error())
		}
		if err != nil {
			log15.Error(errors.Cause(err).Error(), "err", fmt.Sprintf("cause: %s, method: %s, path: %s,  params: %+v, err: %+v", errors.Cause(err), method, r.URL.Path, params, err))
		}
	}()
	path := idRegexp.ReplaceAllString(strings.TrimSuffix(r.URL.Path, "/"), "/{id}$1")
	method = Method(r.Method)
	route, ok := h.routeMap[Path{path, method}]
	if !ok {
		log15.Error("Failed to exec", "path", path, "method", method)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	f := route.Func
	validate := validator.New()
	args := []reflect.Value{
		reflect.ValueOf(ctx),
	}

	funcType := reflect.TypeOf(f)
	if 1 < funcType.NumIn() {
		inputType := funcType.In(1)
		if inputType.Kind() != reflect.Slice {
			if err = validate.Struct(params); err != nil {
				http.Error(w, http.StatusText(400), 400)
				return
			}
		}
		args = append(args, reflect.Indirect(reflect.ValueOf(params)))
	}

	fv := reflect.ValueOf(f)
	results := fv.Call(args)
	var responseJSON []byte
	switch len(results) {
	case 2:
		normalResult, errResult := results[0], results[1]
		if errResult.Interface() == nil {
			if responseJSON, err = json.Marshal(normalResult.Interface()); err != nil {
				http.Error(w, http.StatusText(400), 400)
				return
			}
		} else {
			err = errResult.Interface().(error)
			var code, msg string
			switch v := errors.Cause(err).(type) {
			case errof.UserErr:
				code = string(v)
				msg = v.Error()
			case errof.InternalErr:
				code = string(v)
				msg = v.Error()
			default:
				msg = v.Error()
			}
			returnErr := Error{
				Code:    code,
				Message: msg,
			}
			w.WriteHeader(http.StatusNotAcceptable)
			var _err error
			if responseJSON, _err = json.Marshal(returnErr); _err != nil {
				http.Error(w, http.StatusText(400), 400)
				log15.Error("Failed to marshal err message", "err", _err)
				return
			}
		}
	default:
		http.Error(w, http.StatusText(400), 400)
		log15.Error("Invalid result", "result", results)
		return
	}

	// response to front
	w.Header().Set("Content-Type", "application/json")
	if _, _err := w.Write(responseJSON); _err != nil {
		http.Error(w, http.StatusText(400), 400)
		log15.Error("Failed to write response", "err", _err)
		return
	}
}

func fillIDIntoCtx(next http.Handler) http.Handler {
	var err error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "id")
		var id int64
		if id, err = strconv.ParseInt(rawID, 10, 64); err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
