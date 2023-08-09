package middleware

import (
	"net/http"

	"pular.server/route/process"
)

func ValidationMethods(fn http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		methodAllowed := false
		for _, method := range methods {
			if r.Method == method {
				methodAllowed = true
				break
			}
		}

		if !methodAllowed {
			process.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
			return
		}

		fn(w, r)
	}
}
