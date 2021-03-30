package handler

import (
	"net/http"

	"github.com/soyantonio-w/academy-go-q12021/api/presenter"
	"github.com/soyantonio-w/academy-go-q12021/usecase/launch"

	"github.com/gorilla/mux"
)

// Add comment to exported functions
func SyncLaunches(s *launch.Service, serviceOfData *launch.Service) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := s.SyncLaunches(serviceOfData)

		if err != nil {
			writer.WriteHeader(http.StatusConflict)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

func ListLaunches(s *launch.Service) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		launches, _ := s.ListLaunches()
		// Manage errors

		var presenters []presenter.LaunchPresenter
		for _, l := range launches {
			p := presenter.NewLaunchPresenter(l)
			presenters = append(presenters, p)
		}

		_, _ = writer.Write(presenter.FormatMany(presenters))
	}
}

func GetLaunch(s *launch.Service) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		l, err := s.GetLaunch(vars["id"])

		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

		p := presenter.NewLaunchPresenter(l)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(p.Format())
	}
}
