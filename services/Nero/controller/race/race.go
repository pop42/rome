package race

import (
	"log"
	"net/http"

	"github.com/bluefoxcode/rome/services/Nero/lib/router"
	"github.com/bluefoxcode/rome/services/Nero/lib/util"
	"github.com/bluefoxcode/rome/services/Nero/model/race"
	"github.com/unrolled/render"
)

var (
	url = "/race"
	r   *render.Render
)

type apiError struct {
	Message string
	Code    int
}

// Load the routes.
func Load() {
	log.Println("I am here getting my routes loaded....")
	r = render.New(render.Options{
		IndentJSON: true,
	})
	router.Get(url, Index)
	router.Get(url+"/wipeout", Wipeout)
	// router.Get(url+"/{id}", Show)
}

// Index displays list of heroes
func Index(w http.ResponseWriter, req *http.Request) {
	c := util.Context(w, req)
	items, _, err := race.List(c.DB)

	if err != nil {
		log.Println("Error in controller/race/Index: ", err)
		r.JSON(w, http.StatusBadRequest, apiError{Message: "Bad Request", Code: http.StatusBadRequest})
		return
	}
	if items == nil {
		r.JSON(w, http.StatusOK, map[string]string{})
		return
	}
	r.JSON(w, http.StatusOK, items)

}

// Wipeout will wipeout the race table
func Wipeout(w http.ResponseWriter, req *http.Request) {
	c := util.Context(w, req)
	err := race.Wipeout(c.DB)

	if err != nil {
		log.Println("Error wiping out db: ", err)
		r.JSON(w, http.StatusInternalServerError, apiError{"Internal Server Error", http.StatusInternalServerError})
	}

	r.JSON(w, http.StatusOK, map[string]string{"Message": "Table is empty now"})
}

// // Show displays a single item.
// func Show(w http.ResponseWriter, req *http.Request) {
// 	c := util.Context(w, req)

// 	item, _, err := hero.ByID(c.DB, c.Param("id"))

// 	if err != nil {
// 		r.JSON(w, http.StatusNotFound, apiError{Message: "Not Found", Code: http.StatusNotFound})
// 		return
// 	}

// 	r.JSON(w, http.StatusOK, item)

// }
