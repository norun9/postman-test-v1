package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/norun9/postmantest/internal/api/usecase/input"
)

func (h restHandler) GetPostRouter(router chi.Router) {
	router.Route("/{id}", func(router chi.Router) {
		router.Use(fillIDIntoCtx)
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			params := input.GetPost{
				ID: ctx.Value("id").(int64),
			}
			h.Exec(ctx, w, r, params)
		})
	})
}
