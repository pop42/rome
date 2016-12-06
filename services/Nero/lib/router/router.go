package router

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	r         *mux.Router
	infoMutex sync.RWMutex
)

// init sets up the router.
func init() {
	ResetConfig()
}

// ResetConfig creates a new instance.
func ResetConfig() {
	infoMutex.Lock()
	r = mux.NewRouter()
	infoMutex.Unlock()
}

// Instance returns the router.
func Instance() *mux.Router {
	infoMutex.RLock()
	defer infoMutex.RUnlock()
	return r
}

// Param returns the URL parameter.
func Param(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}
