package launch

import (
	"fmt"
	"github.com/soyantonio-w/academy-go-q12021/entity"
	"strconv"
)

type Service struct {
	repo entity.LaunchRepo
}

func NewService(r entity.LaunchRepo) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetLaunch(launchId string) (entity.Launch, error) {
	id, err := strconv.Atoi(launchId)
	if err != nil {
		return entity.Launch{}, fmt.Errorf("non valid id for launch")
	}
	launch, err := s.repo.Get(entity.LaunchId(id))
	return launch, err
}

func (s *Service) ListLaunches() ([]entity.Launch, error) {
	launches, err := s.repo.GetLaunches()
	if err != nil {
		return []entity.Launch{}, err
	}
	return launches, nil
}
