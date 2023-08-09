package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"pular.server/database"
	"pular.server/route"
	"pular.server/route/middleware"
	"pular.server/route/process"
)

func Routes() []route.Handler {
	return []route.Handler{
		{
			Path:        "/login",
			Methods:     []string{"POST"},
			HandlerFunc: LoginHandle,
		},
		{
			Path:        "/info",
			Methods:     []string{"GET"},
			HandlerFunc: InfoHandle,
			Middlewares: []middleware.Middleware{
				ValidationRequest,
			},
		},
	}
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginReq

	err := process.ReadJsonBody(w, r.Body, &reqBody)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	auth, err := Login(db, reqBody)

	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, auth)
}

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sql.DB)
	identity := r.Context().Value("identity").(*Identity)

	fmt.Println("identity :", identity.Account_id)

	account, err := database.GetAccountInfoById(db, identity.Account_id)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, account)
}
