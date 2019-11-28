package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"../models"
	"../utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{domain}", GetSite)
	return router
}

func GetSite(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	site, err := models.FetchSite(domain)

	if err == sql.ErrNoRows {
		utils.APIInfo(domain)

	} else if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, site)
}
