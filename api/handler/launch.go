package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/presenter"
	"github.com/soyantonio-w/academy-go-q12021/infrastructure/repository/csv"
	"net/http"
	"strconv"
)


func ListLaunches(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	r := csv.NewRepository()
	launches, _ := r.GetLaunches()

	var presenters []presenter.LaunchPresenter
	for _, launch := range launches {
		p := presenter.NewLaunchPresenter(launch)
		presenters = append(presenters, p)
	}

	writer.Write(presenter.FormatMany(presenters))
}

func GetLaunch(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	launchId, _ := strconv.Atoi(vars["id"])

	r := csv.NewRepository()
	launches, _ := r.GetLaunches()

	var presenters []presenter.LaunchPresenter
	for _, launch := range launches {
		p := presenter.NewLaunchPresenter(launch)
		presenters = append(presenters, p)
	}

	fmt.Println(launchId)
	writer.Write(presenter.FormatMany(presenters))
}
