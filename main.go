package main

import (
	"log"
	"net/http"
	"os"

	"github.com/soyantonio-w/academy-go-q12021/api/config"
	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/gspacex"
	"github.com/soyantonio-w/academy-go-q12021/usecase"

	"github.com/gorilla/mux"
)

const (
	ExitCouldNotLoadConfigFile = iota + 1
)

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Printf("could not load config %v", err)
		os.Exit(ExitCouldNotLoadConfigFile)
	}

	cacheLaunchService := usecase.NewService(csv.NewRepository())
	flightLaunchService := usecase.NewService(gspacex.NewRepository(cfg.GetSpacexAddress()))

	r := mux.NewRouter()
	r.HandleFunc("/cache/launches", handler.ListLaunches(cacheLaunchService))
	r.HandleFunc("/cache/sync", handler.SyncLaunches(cacheLaunchService, flightLaunchService))
	r.HandleFunc("/cache/launch/{id:[0-9]+}", handler.GetLaunch(cacheLaunchService))

	r.HandleFunc("/flight/launches", handler.ListLaunches(flightLaunchService))
	r.HandleFunc("/flight/launch/{id:[0-9]+}", handler.GetLaunch(flightLaunchService))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(cfg.GetAppAddress(), nil))
}
