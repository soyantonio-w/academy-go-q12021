package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - expected methods to be linked with the router
type Handler interface {
	GetLaunch(w http.ResponseWriter, r *http.Request)
	ListLaunches(w http.ResponseWriter, r *http.Request)
	SyncLaunches(w http.ResponseWriter, r *http.Request)
	FilterLaunches(w http.ResponseWriter, r *http.Request)
}

// New - creates a mux router
func New(flight Handler, cache Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cache/launch/{id:[0-9]+}", cache.GetLaunch).Methods("GET")
	r.HandleFunc("/cache/launches/sync", cache.SyncLaunches).Methods("GET")
	r.HandleFunc("/cache/launches/list", cache.ListLaunches).Methods("GET")
	r.HandleFunc("/cache/launches/filter", cache.FilterLaunches).
		Queries(
			"type", "{type:odd|even}",
			"items", "{items:[0-9]+}",
			"items_per_workers", "{items_per_workers:[0-9]+}",
		).
		Methods("GET")

	r.HandleFunc("/flight/launch/{id:[0-9]+}", flight.GetLaunch).Methods("GET")
	r.HandleFunc("/flight/launches/list", flight.ListLaunches).Methods("GET")

	return r
}
