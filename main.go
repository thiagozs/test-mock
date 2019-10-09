package main

import (
	"fmt"

	"github.com/thiagozs/test-mock/dbs"
)

func main() {
	fmt.Println("Hello Redis Mock")

	redis := dbs.NewRedis()

	if err := redis.Ping(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")
}
