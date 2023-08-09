package process

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pular.server/route/status"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	dataJson, err := json.Marshal(data)

	if err != nil {
		WriteInternal(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
}

func WriteSuccess(w http.ResponseWriter) {
	WriteJson(w, status.Success())
}

func WriteInternal(w http.ResponseWriter, err error) {
	fmt.Println("process error: ", err)
	WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	dataJson, _ := json.Marshal(status.Message{
		Message: message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dataJson)
}
