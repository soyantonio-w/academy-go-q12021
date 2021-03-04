package handler

import (
	"github.com/gorilla/mux"
	"github.com/soyantonio-w/academy-go-q12021/api/presenter"
	"github.com/soyantonio-w/academy-go-q12021/entity"
	"net/http"
	"strconv"
)

func ListLaunches(repo entity.LaunchRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		launches, _ := repo.GetLaunches()

		var presenters []presenter.LaunchPresenter
		for _, launch := range launches {
			p := presenter.NewLaunchPresenter(launch)
			presenters = append(presenters, p)
		}

		_, _ = writer.Write(presenter.FormatMany(presenters))
	}
}

func GetLaunch(repo entity.LaunchRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		launchId, _ := strconv.Atoi(vars["id"])

		launch, err := repo.Get(entity.LaunchId(launchId))
		if err != nil {
			writer.WriteHeader(404)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

		p := presenter.NewLaunchPresenter(launch)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(p.Format())
	}
}
