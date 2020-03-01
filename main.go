package main

import (
	"log"

	"github.com/hlscalon/go-react-boilerplate/router"
)

func main() {
	r, err := router.New()
	if err != nil {
		log.Fatal(err)
	}

	r.Init("3000")
}
