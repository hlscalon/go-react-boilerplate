//
// This program is a small example of a boilerplate that embraces go on server-side and react on client-side
// On server-side it is divided in some packages:
// 	- router
// 	- controllers
// 	- models
// 	- utils
//
// It uses chi for routing and upper.io as database layer
//

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
