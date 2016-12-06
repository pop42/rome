package controller

import (
	"log"

	"github.com/bluefoxcode/rome/services/Nero/controller/hero"
	"github.com/bluefoxcode/rome/services/Nero/controller/race"
)

// LoadRoutes loads the routes for each of the controllers
func LoadRoutes() {
	hero.Load()
	log.Println("Here, loading routes....")
	race.Load()
}
