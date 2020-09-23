package server

import (
	"github.com/gorilla/mux"
)

// NewRouter function get the gorilla mux router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addMainHandler(r)
	return r
}

func addMainHandler(r *mux.Router) {
	r.HandleFunc("/flatten", GetFlattenArray).Methods("POST")
}
