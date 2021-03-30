package usecase

import (
	"fmt"
	"strconv"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

// Service - holds the launch use case
type Service struct {
	repo entity.LaunchRepo
}

// NewService - creates a launch use case
func NewService(r entity.LaunchRepo) *Service {
	return &Service{
		repo: r,
	}
}

// GetLaunch - returns a Launch that matches with the launchId
func (s *Service) GetLaunch(launchId string) (entity.Launch, error) {
	id, err := strconv.Atoi(launchId)
	if err != nil {
		return entity.Launch{}, fmt.Errorf("non valid id for launch")
	}
	launch, err := s.repo.Get(entity.LaunchId(id))
	return launch, err
}

// ListLaunches - returns all the launches
func (s *Service) ListLaunches() ([]entity.Launch, error) {
	launches, err := s.repo.GetLaunches()
	if err != nil {
		return []entity.Launch{}, err
	}
	return launches, nil
}

// ListLaunches - loads all the launches from the provided launch use case
func (s *Service) SyncLaunches(data *Service) error {
	launches, err := data.ListLaunches()

	if err != nil {
		return fmt.Errorf("could not sync launches")
	}

	err = s.repo.SyncAll(launches)

	if err != nil {
		return fmt.Errorf("could not sync launches")
	}

	return nil
}
