package usecase

import "math"

type interval struct {
	iteration,
	items,
	length int
}

func (i *interval) getIntervals() (int, int) {
	return i.getStartInterval(), i.getEndInterval()
}

func (i *interval) getEndInterval() int {
	// The iteration should at zero
	nextIteration := i.iteration + 1
	return int(math.Min(float64(nextIteration*i.items), float64(i.length)))
}

func (i *interval) getStartInterval() int {
	return int(math.Min(float64(i.iteration*i.items), float64(i.length)))
}
