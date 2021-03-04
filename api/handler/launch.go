package handler

import (
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/presenter"
	"github.com/soyantonio-w/academy-go-q12021/entity"
	"net/http"
	"strconv"
)

func GetLaunch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	launchId, _ := strconv.Atoi(vars["id"])
	p := presenter.LaunchPresenter{ ID: entity.LaunchId(launchId) }

	writer.Write(p.Format())
	writer.Header().Set("Content-Type", "application/json")
}
