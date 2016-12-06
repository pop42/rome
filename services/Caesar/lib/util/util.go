package util

import (
	"net/http"
	"sync"

	"github.com/bluefoxcode/rome/services/Caesar/lib/router"
	"github.com/jmoiron/sqlx"
)

var (
	dbInfo *sqlx.DB
	mutex  sync.RWMutex
)

// StoreDB stores the database connection settigns so controller functions can access them safely.
func StoreDB(db *sqlx.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

// Info structures the application settings.
type Info struct {
	W  http.ResponseWriter
	R  *http.Request
	DB *sqlx.DB
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {
	mutex.RLock()
	i := Info{
		W:  w,
		R:  r,
		DB: dbInfo,
	}
	mutex.RUnlock()
	return i
}

// CheckErr is quick utility for panicking
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Param gets the URL parameter.
func (c *Info) Param(name string) string {
	return router.Param(c.R, name)
}
