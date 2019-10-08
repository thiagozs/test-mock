package main

import (
	"fmt"

	"github.com/thiagozs/test-mock/dbs"
)

func main() {
	fmt.Println("Hello")

	_, err := dbs.NewRedis()
	if err != nil {
		panic(err)
	}
}
