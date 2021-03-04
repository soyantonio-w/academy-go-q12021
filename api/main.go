package main

import (
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"github.com/soyantonio-w/academy-go-q12021/usecase/launch"
	"log"
	"net/http"
)

const httpAddr = "localhost:8080"

func main() {
	launchRepo := csv.NewRepository()
	launchService := launch.NewService(launchRepo)

	r := mux.NewRouter()
	r.HandleFunc("/launches", handler.ListLaunches(launchService))
	r.HandleFunc("/launch/{id:[0-9]+}", handler.GetLaunch(launchService))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
