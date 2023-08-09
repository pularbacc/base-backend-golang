package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
}

func (r Request) Send(dataInp []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(dataInp))
	if err != nil {
		return nil, err
	}

	// Set additional headers from the config
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r Request) SendStruct(dataInp interface{}, dataOut interface{}) error {
	inpByte, err := json.Marshal(dataInp)
	if err != nil {
		return err
	}

	outByte, err := r.Send(inpByte)
	if err != nil {
		return err
	}

	err = json.Unmarshal(outByte, &dataOut)
	if err != nil {
		return err
	}
	return nil
}

func (r Request) SendString(dataInp string) (string, error) {
	inpByte := []byte(dataInp)

	outByte, err := r.Send(inpByte)
	if err != nil {
		return "", err
	}

	outString := string(outByte)

	return outString, nil
}
