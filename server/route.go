package server

import (
	"fmt"
	"net/http"

	"pular.server/route"
	"pular.server/route/process"
)

func Routes() []route.Handler {
	return []route.Handler{
		{
			Path:        "",
			Methods:     []string{"GET"},
			HandlerFunc: IndexHandle,
		},
	}
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("on index handle")
	process.WriteSuccess(w)
}
