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

	csvLaunchRepo := csv.NewRepository()
	spacexLaunchRepo := gspacex.NewRepository(cfg.GetSpacexAddress())

	csvLaunchUseCase := usecase.LaunchNew(csvLaunchRepo)
	spacexLaunchUseCase := usecase.LaunchNew(spacexLaunchRepo)

	r := mux.NewRouter()
	r.HandleFunc("/cache/launches", handler.ListLaunches(csvLaunchUseCase))
	r.HandleFunc("/cache/sync", handler.SyncLaunches(csvLaunchUseCase, spacexLaunchUseCase))
	r.HandleFunc("/cache/launch/{id:[0-9]+}", handler.GetLaunch(csvLaunchUseCase))

	r.HandleFunc("/flight/launches", handler.ListLaunches(spacexLaunchUseCase))
	r.HandleFunc("/flight/launch/{id:[0-9]+}", handler.GetLaunch(spacexLaunchUseCase))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(cfg.GetAppAddress(), nil))
}
