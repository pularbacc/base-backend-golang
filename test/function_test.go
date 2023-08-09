package test

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	hello("haha")
}

func hello(mess ...string) {
	fmt.Println(mess)
}
