package middleware

import (
	"fmt"
	"net/http"
)

func Log(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s]:%s\n", r.Method, r.URL.Path)
		fn(w, r)
	}
}
