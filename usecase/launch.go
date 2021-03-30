package usecase

import (
	"fmt"
	"strconv"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

// LaunchUseCase - holds the launch use case
type LaunchUseCase struct {
	repo entity.LaunchRepo
}

// NewService - creates a launch use case
func NewService(r entity.LaunchRepo) *LaunchUseCase {
	return &LaunchUseCase{
		repo: r,
	}
}

// GetLaunch - returns a Launch that matches with the launchId
func (s *LaunchUseCase) GetLaunch(launchId string) (entity.Launch, error) {
	id, err := strconv.Atoi(launchId)
	if err != nil {
		return entity.Launch{}, fmt.Errorf("non valid id for launch")
	}
	launch, err := s.repo.Get(entity.LaunchId(id))
	return launch, err
}

// ListLaunches - returns all the launches
func (s *LaunchUseCase) ListLaunches() ([]entity.Launch, error) {
	launches, err := s.repo.GetLaunches()
	if err != nil {
		return []entity.Launch{}, err
	}
	return launches, nil
}

// ListLaunches - loads all the launches from the provided launch use case
func (s *LaunchUseCase) SyncLaunches(data *LaunchUseCase) error {
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
