package usecase

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

// LaunchUseCase - holds the launch use case
type LaunchUseCase struct {
	repo entity.LaunchRepo
}

// LaunchNew - creates a launch use case
func LaunchNew(r entity.LaunchRepo) *LaunchUseCase {
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

// FilterLaunches - returns all the launches that matches with the query
func (s *LaunchUseCase) FilterLaunches(filterType string, items int, itemsPerWorker int) ([]entity.Launch, error) {
	launches, err := s.repo.GetLaunches()
	if err != nil {
		return []entity.Launch{}, err
	}
	if itemsPerWorker <= 0 {
		return []entity.Launch{}, fmt.Errorf("items per worker should be greater than 0")
	}

	totalLaunches := len(launches)
	t := NewLaunchFilterType(filterType)
	inRangeItems := int(math.Min(float64(items), float64(totalLaunches)))
	workers := calculateRequiredWorkers(totalLaunches, itemsPerWorker)
	results := make(chan entity.Launch)
	done := make(chan int, workers)

	var wg sync.WaitGroup
	wg.Add(workers)
	shutdown := make(chan struct{})

	for w := 0; w < workers; w++ {
		i := &interval{
			iteration: w,
			items:     itemsPerWorker,
			length:    totalLaunches,
		}
		start, end := i.getIntervals()
		go worker(&wg, done, shutdown, w, launches[start:end], results, t)
	}

	filteredLaunches := filterLaunches(results, done, workers, inRangeItems)

	log.Printf("wating %d", workers)
	close(shutdown)
	wg.Wait()
	log.Printf("done")

	return filteredLaunches, nil
}

func filterLaunches(results chan entity.Launch, done chan int, workers, items int) []entity.Launch {
	workersDone := 0
	filteredLaunches := make([]entity.Launch, 0, items)
	if workers == 0 || items == 0 {
		return filteredLaunches
	}
	for {
		select {
		case launch := <-results:
			filteredLaunches = append(filteredLaunches, launch)

			if len(filteredLaunches) == items {
				return filteredLaunches
			}
		case <-done:
			workersDone += 1
			if workersDone == workers {
				return filteredLaunches
			}
		}
	}
}

func calculateRequiredWorkers(items int, itemsPerWorker int) int {
	return int(math.Ceil(float64(items) / float64(itemsPerWorker)))
}

func worker(
	wg *sync.WaitGroup,
	workerDone chan int,
	shutdown chan struct{},
	id int, launches []entity.Launch,
	results chan<- entity.Launch,
	filterType *LaunchFilterType,
) {
	log.Printf("worker %d started job", id)

	filteredLaunches := filterLaunchesByType(launches, *filterType)
	if len(filteredLaunches) == 0 {
		log.Printf("worker %d finished job", id)
		workerDone <- id
		wg.Done()
		return
	}

	i := 0
	for {
		select {
		case results <- filteredLaunches[i]:
			i++
			if i == len(filteredLaunches) {
				log.Printf("worker %d finished job", id)
				workerDone <- id
				wg.Done()
				return
			}
		case <-shutdown:
			workerDone <- id
			wg.Done()
			log.Printf("worker %d finished job wii", id)
			return
		}
	}
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
