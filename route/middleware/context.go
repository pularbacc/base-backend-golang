package middleware

import (
	"context"
	"database/sql"
	"net/http"
)

func MakeContext(fn http.HandlerFunc, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "db", db)

		fn(w, r.WithContext(ctx))
	}
}
