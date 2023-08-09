package auto

import (
	"database/sql"
	"net/http"
	"strconv"

	"pular.server/database"
	"pular.server/route"
	"pular.server/route/process"
)

func Routes() []route.Handler {
	return []route.Handler{
		{
			Path:        "/create",
			Methods:     []string{"POST"},
			HandlerFunc: CreateHandle,
		},
		{
			Path:        "/update",
			Methods:     []string{"PATCH"},
			HandlerFunc: UpdateHandle,
		},
		{
			Path:        "/delete",
			Methods:     []string{"DELETE"},
			HandlerFunc: DeleteHandle,
		},
		{
			Path:        "/detail",
			Methods:     []string{"GET"},
			HandlerFunc: DetailHandle,
		},
		{
			Path:        "",
			Methods:     []string{"GET"},
			HandlerFunc: GetListHandle,
		},
		{
			Path:        "/run",
			Methods:     []string{"POST"},
			HandlerFunc: RunHandle,
		},
	}
}

func CreateHandle(w http.ResponseWriter, r *http.Request) {
	var reqBody database.Auto
	err := process.ReadJsonBody(w, r.Body, &reqBody)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	db := r.Context().Value("db").(*sql.DB)

	auto, err := database.CreateAuto(db, reqBody)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, auto)
}

func UpdateHandle(w http.ResponseWriter, r *http.Request) {
	var reqBody database.Auto
	err := process.ReadJsonBody(w, r.Body, &reqBody)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	db := r.Context().Value("db").(*sql.DB)

	err = database.UpdateAuto(db, reqBody)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteSuccess(w)
}

func DeleteHandle(w http.ResponseWriter, r *http.Request) {
	idS := process.ReadQuery(r, "id")

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	err = database.DeleteAuto(db, id)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteSuccess(w)
}

func DetailHandle(w http.ResponseWriter, r *http.Request) {
	idS := process.ReadQuery(r, "id")

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	auto, err := database.GetAutoByID(db, id)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, auto)
}

func GetListHandle(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sql.DB)

	autos, err := database.GetAutoList(db)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, autos)
}

func RunHandle(w http.ResponseWriter, r *http.Request) {
	idS := process.ReadQuery(r, "id")

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		process.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	auto, err := database.GetAutoByID(db, id)
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	authHeader := r.Context().Value("Authorization").(string)

	results, err := Run(auto, RunConfig{
		Authorization: authHeader,
	})
	if err != nil {
		process.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	process.WriteJson(w, results)
}
