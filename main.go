package main

import (
	"log"
	"net/http"

	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/gspacex"
	"github.com/soyantonio-w/academy-go-q12021/usecase/launch"

	"github.com/gorilla/mux"
)

// TODO move this to config file - .conf (viper)
const httpAddr = "localhost:8080"
const spacexAPI = "https://api.spacex.land/graphql/"

func main() {
	// cfg := config.Load()

	cacheLaunchService := launch.NewService(csv.NewRepository())
	flightLaunchService := launch.NewService(gspacex.NewRepository(spacexAPI))

	r := mux.NewRouter()
	r.HandleFunc("/cache/launches", handler.ListLaunches(cacheLaunchService))
	r.HandleFunc("/cache/sync", handler.SyncLaunches(cacheLaunchService, flightLaunchService))
	r.HandleFunc("/cache/launch/{id:[0-9]+}", handler.GetLaunch(cacheLaunchService))

	r.HandleFunc("/flight/launches", handler.ListLaunches(flightLaunchService))
	r.HandleFunc("/flight/launch/{id:[0-9]+}", handler.GetLaunch(flightLaunchService))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
