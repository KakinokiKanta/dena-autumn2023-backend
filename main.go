package main

import (
	"fmt"

	"github.com/shinbunbun/go-gin-template/router"
)

func main() {
	r := router.Route()
	err := r.Run("localhost:8081")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
}
