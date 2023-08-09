package client

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestSendWithStruct(t *testing.T) {
	request := Request{
		Method: "POST",
		Url:    "https://test.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	inpStruct := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Service  string `json:"service"`
	}{
		Email:    "test@gmail.com",
		Password: "123456Aa",
		Service:  "TEAM",
	}

	inpByte, err := json.Marshal(inpStruct)
	if err != nil {
		log.Fatal(err)
	}

	outByte, err := request.Send(inpByte)
	if err != nil {
		log.Fatal(err)
	}

	var outStruct struct {
		Status string `json:"status"`
		Data   struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	err = json.Unmarshal(outByte, &outStruct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dataOut :", outStruct)
}

func TestSendStruct(t *testing.T) {
	request := Request{
		Method: "POST",
		Url:    "https://test.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	inpStruct := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Service  string `json:"service"`
	}{
		Email:    "test@gmail.com",
		Password: "123456Aa",
		Service:  "TEAM",
	}

	var outStruct struct {
		Status string `json:"status"`
		Data   struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	err := request.SendStruct(inpStruct, &outStruct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(outStruct)
}

func TestSendWithByte(t *testing.T) {
	request := Request{
		Method: "POST",
		Url:    "https://test.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	inpString := "{\"email\":\"test@gmail.com\",\"password\":\"123456Aa\",\"service\":\"TEAM\"}"

	inpByte := []byte(inpString)

	outByte, err := request.Send(inpByte)
	if err != nil {
		log.Fatal(err)
	}

	outString := string(outByte)

	fmt.Println("dataOut :", outString)
}

func TestSendString(t *testing.T) {
	request := Request{
		Method: "POST",
		Url:    "https://test.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	inpString := "{\"email\":\"test@gmail.com\",\"password\":\"123456Aa\",\"service\":\"TEAM\"}"

	outString, err := request.SendString(inpString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(outString)
}
