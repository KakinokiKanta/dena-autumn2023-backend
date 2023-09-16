package main

import (
	"fmt"

	"github.com/shinbunbun/dena-autumn-backend/server/router"
)

func main() {
	r := router.Route()
	err := r.Run("0.0.0.0:8081")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
}
