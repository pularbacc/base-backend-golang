package process

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJsonBody(w http.ResponseWriter, r io.Reader, data interface{}) error {
	err := json.NewDecoder(r).Decode(&data)

	if err != nil {
		return err
	}

	return nil
}

func ReadQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
