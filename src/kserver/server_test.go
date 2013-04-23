package kserver

import (
	"fmt"
	"testing"
)

func TestinitServer(t *testing.T) {
	initServer(":8080")
	fmt.Println("----")
}
