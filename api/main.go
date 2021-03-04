package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const httpAddr = ":8082"

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/launch/{id:[0-9]+}", LaunchHandler)

	http.Handle("/", r)
	http.ListenAndServe(httpAddr, nil)
}


