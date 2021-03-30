package usecase

import "github.com/soyantonio-w/academy-go-q12021/entity"

func removeOddLaunches(launches []entity.Launch) []entity.Launch {
	filtered := make([]entity.Launch, 0)
	for i := range launches {
		if launches[i].LaunchId%2 == 1 {
			continue
		}
		filtered = append(filtered, launches[i])
	}
	return filtered
}

func removeEvenLaunches(launches []entity.Launch) []entity.Launch {
	filtered := make([]entity.Launch, 0)
	for i := range launches {
		if launches[i].LaunchId%2 == 0 {
			continue
		}
		filtered = append(filtered, launches[i])
	}
	return filtered
}
