package route

import (
	"database/sql"
	"net/http"

	"pular.server/route/middleware"
)

type Handler struct {
	Path        string
	Methods     []string
	Middlewares []middleware.Middleware
	HandlerFunc http.HandlerFunc
	Group       []Handler
}

func (h *Handler) Make(db *sql.DB) http.HandlerFunc {
	handle := h.HandlerFunc

	// middleware by request
	handle = middleware.Chain(h.Middlewares...)(handle)

	// middleware global
	handle = middleware.ValidationMethods(handle, h.Methods)
	handle = middleware.CORS(handle)
	handle = middleware.Log(handle)

	handle = middleware.MakeContext(handle, db)

	return handle
}

type GroupHandler struct {
	Path        string
	Handlers    []Handler
	Middlewares []middleware.Middleware
}

func (g *GroupHandler) Make(mux *http.ServeMux, db *sql.DB) {
	for _, handler := range g.Handlers {
		if handler.Group != nil {
			groups := GroupHandler{
				Path:        g.Path + handler.Path,
				Handlers:    handler.Group,
				Middlewares: append(g.Middlewares, handler.Middlewares...),
			}
			groups.Make(mux, db)
		} else {
			handler.Middlewares = append(g.Middlewares, handler.Middlewares...)
			mux.HandleFunc(g.Path+handler.Path, handler.Make(db))
		}
	}
}
