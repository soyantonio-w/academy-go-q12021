package main

import (
	"log"
	"net/http"
	"os"

	"github.com/soyantonio-w/academy-go-q12021/api/config"
	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/gspacex"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/router"
	"github.com/soyantonio-w/academy-go-q12021/usecase"
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

	flightHandler := handler.New(spacexLaunchUseCase, nil)
	cacheHandler := handler.New(csvLaunchUseCase, spacexLaunchUseCase)
	r := router.New(flightHandler, cacheHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(cfg.GetAppAddress(), nil))
}
