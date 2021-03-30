package usecase

import (
	"github.com/soyantonio-w/academy-go-q12021/entity"
)

// LaunchFilterType - type used to hold the available values of launch type filter
type LaunchFilterType int

const (
	// Even - default value of the filter, only includes launches with even id
	Even LaunchFilterType = iota
	// Odd - only includes launches with odd id
	Odd
)

// String - implement stringer
func (t LaunchFilterType) String() string {
	return toString[t]
}

// NewLaunchFilterType - creates a filter type from string.
// If the string is not valid then it gets converted to "even"
func NewLaunchFilterType(t string) *LaunchFilterType {
	launchFilterType := toId[t]
	return &launchFilterType
}

var toString = map[LaunchFilterType]string{
	Odd:  "odd",
	Even: "even",
}

var toId = map[string]LaunchFilterType{
	"odd":  Odd,
	"even": Even,
}

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

func filterLaunchesByType(launches []entity.Launch, filterType LaunchFilterType) []entity.Launch {
	if filterType == Odd {
		return removeEvenLaunches(launches)
	}
	return removeOddLaunches(launches)
}
