package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/norun9/postmantest/internal/api/injector"
	"github.com/norun9/postmantest/pkg/config"
)

func main() {
	c := config.Prepare()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   c.HTTP.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	})

	handler := injector.InitializeRestHandler(c.HTTP, c.MySQL)
	router.Route("/", handler.GetRoute)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.HTTP.Port), router))

	//  curl -XGET -H 'Content-Type: application/json; charset=utf-8' "http://localhost:7070/v1/posts/2" | jq .
	// {
	// 		"id": 2,
	// 		"content": "fugafuga",
	// 		"created_at": "2022-08-08T20:00:00+09:00",
	// 		"updated_at": "2022-08-08T20:00:00+09:00"
	// }
}
