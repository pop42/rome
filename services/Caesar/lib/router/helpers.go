package router

import "net/http"

// Get is a shortcut for HandleFunc(path, handle).Methods("GET")
func Get(path string, fn http.HandlerFunc) {
	infoMutex.Lock()
	r.HandleFunc(path, fn).Methods("GET")
	infoMutex.Unlock()
}

// Post is a shortcut for HandleFunc(path, handle).Methods("POST")
func Post(path string, fn http.HandlerFunc) {
	infoMutex.Lock()
	r.HandleFunc(path, fn).Methods("POST")
	infoMutex.Unlock()
}
