package main

import (
	"log"

	"github.com/hlscalon/go-react-boilerplate/router"
	"github.com/hlscalon/go-react-boilerplate/models"
)

func main() {
	db, err := models.NewDB("localhost", "go_react_boilerplate", "root", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.Init(db, "3000")
}
