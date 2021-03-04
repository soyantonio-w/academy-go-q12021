package main

import (
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/handler"
	"net/http"
)

const httpAddr = ":8080"

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/launch/{id:[0-9]+}", handler.GetLaunch)

	http.Handle("/", r)
	http.ListenAndServe(httpAddr, nil)
}


