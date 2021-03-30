package handler

import (
	"net/http"

	"github.com/soyantonio-w/academy-go-q12021/api/presenter"
	"github.com/soyantonio-w/academy-go-q12021/usecase"

	"github.com/gorilla/mux"
)

// Handler - struct
type Handler struct {
	useCase, s *usecase.LaunchUseCase
}

// New - returns a handler implementation
func New(useCase *usecase.LaunchUseCase, s *usecase.LaunchUseCase) *Handler {
	return &Handler{useCase, s}
}

// SyncLaunches - provides the logic to sync launches and return and http response
func (h *Handler) SyncLaunches(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := h.useCase.SyncLaunches(h.s)

	if err != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

// ListLaunches - provides the logic to list all launches and return them as a http response
func (h *Handler) ListLaunches(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	launches, err := h.useCase.ListLaunches()

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	var presenters []presenter.LaunchPresenter
	for _, l := range launches {
		p := presenter.NewLaunchPresenter(l)
		presenters = append(presenters, p)
	}

	_, _ = writer.Write(presenter.FormatMany(presenters))
}

// GetLaunch - provides the logic to get only a launch as a http response
func (h *Handler) GetLaunch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	l, err := h.useCase.GetLaunch(vars["id"])

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	p := presenter.NewLaunchPresenter(l)
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(p.Format())
}
