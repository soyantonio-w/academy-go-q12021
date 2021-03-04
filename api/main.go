package main

import (
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	launchRepo := csv.NewRepository()

	r := mux.NewRouter()
	r.HandleFunc("/launches", handler.ListLaunches(launchRepo))
	r.HandleFunc("/launch/{id:[0-9]+}", handler.GetLaunch(launchRepo))

	http.Handle("/", r)
	_ = http.ListenAndServe(httpAddr, nil)
}
