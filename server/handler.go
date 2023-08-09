package server

import (
	"database/sql"
	"net/http"

	"pular.server/auth"
	"pular.server/auto"
	"pular.server/route"
	"pular.server/route/middleware"
)

func Handlers(mux *http.ServeMux, db *sql.DB) {
	indexGroups := route.GroupHandler{
		Path:     "/",
		Handlers: Routes(),
	}
	indexGroups.Make(mux, db)

	authGroups := route.GroupHandler{
		Path:     "/auth",
		Handlers: auth.Routes(),
	}
	authGroups.Make(mux, db)

	autoGroups := route.GroupHandler{
		Path:     "/auto",
		Handlers: auto.Routes(),
		Middlewares: []middleware.Middleware{
			auth.ValidationRequest,
		},
	}
	autoGroups.Make(mux, db)
}
