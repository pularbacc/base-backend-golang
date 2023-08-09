package auto

import (
	"pular.server/client"
	"pular.server/database"
)

type RunConfig struct {
	Authorization string `json:"authorization"`
}

func Run(auto database.Auto, config RunConfig) ([]string, error) {
	var results []string

	for _, cmd := range auto.Cmds {
		request := client.Request{
			Method: cmd.Method,
			Url:    "http://localhost:3051" + cmd.Url,
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": config.Authorization,
			},
		}

		out, err := request.SendString(cmd.Body)
		if err != nil {
			return nil, err
		}

		results = append(results, out)
	}

	return results, nil
}
