package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestJson(t *testing.T) {
	user := User{
		Name: "nguyen",
		Age:  30,
	}

	userJson, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(userJson))
}
