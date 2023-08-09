package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"pular.server/env"
)

func Init(db *sql.DB) {
	fmt.Println("Server Init ... ")

	mux := http.NewServeMux()

	Handlers(mux, db)

	server := &http.Server{
		Addr:    env.PORT,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
}
