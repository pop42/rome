package hero

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bluefoxcode/rome/services/Nero/lib/router"
	"github.com/bluefoxcode/rome/services/Nero/lib/util"
	"github.com/bluefoxcode/rome/services/Nero/model/hero"
	"github.com/unrolled/render"
)

var (
	url = "/hero"
	r   *render.Render
)

type apiError struct {
	Message string
	Code    int
}

// Load the routes.
func Load() {
	r = render.New(render.Options{
		IndentJSON: true,
	})
	router.Get(url, Index)
	router.Post(url, Create)
	router.Get(url+"/{id}", Show)
}

// Index displays list of heroes
func Index(w http.ResponseWriter, req *http.Request) {
	c := util.Context(w, req)
	items, _, err := hero.List(c.DB)

	if err != nil {
		log.Println("Error in controller/hero/Index: ", err)
		r.JSON(w, http.StatusBadRequest, apiError{Message: "Bad Request", Code: http.StatusBadRequest})
		return
	}
	if items == nil {
		r.JSON(w, http.StatusOK, map[string]string{})
		return
	}
	r.JSON(w, http.StatusOK, items)

}

// Show displays a single item.
func Show(w http.ResponseWriter, req *http.Request) {
	c := util.Context(w, req)

	item, _, err := hero.ByID(c.DB, c.Param("id"))

	if err != nil {
		r.JSON(w, http.StatusNotFound, apiError{Message: "Not Found", Code: http.StatusNotFound})
		return
	}

	r.JSON(w, http.StatusOK, item)

}

// Create adds an item.
func Create(w http.ResponseWriter, req *http.Request) {
	c := util.Context(w, req)
	var h hero.Item
	b, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(b, &h)

	_, err := hero.Create(c.DB, h.Name, h.Description)

	if err != nil {
		r.JSON(w, http.StatusInternalServerError, err)
		return
	}

	r.JSON(w, http.StatusCreated, "New Hero Created.")
}
