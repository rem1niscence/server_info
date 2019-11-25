package main

import (
	"log"
	"net/http"

	"./api/controllers"
	"./api/database"
	"./api/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/", func(r chi.Router) {
		r.Mount("/servers", controllers.Routes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	db := database.InitDb()

	// Inject db into db models file
	// TODO: Find better approach to have a global single db connection
	models.Db = db

	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000", router))
}