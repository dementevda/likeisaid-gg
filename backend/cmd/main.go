package main

import (
	"fmt"
	"log"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api"
)

func main() {
	api := api.New()
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(api)
}
