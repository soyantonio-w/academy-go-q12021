package presenter

import (
	"encoding/json"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

// LaunchPresenter - holds the representation of a launch compatible with JSON
type LaunchPresenter struct {
	ID                   entity.LaunchId   `json:"id"`
	LaunchDate           entity.LaunchDate `json:"launch_date"`
	entity.LaunchSuccess `json:"success"`
	entity.LaunchDetails `json:"details"`
	entity.MissionName   `json:"mission_name"`
	entity.RocketName    `json:"rocket_name"`
	entity.VideoLink     `json:"video_link"`
}

// NewLaunchPresenter - creates a launch representation from the Launch entity
func NewLaunchPresenter(launch entity.Launch) LaunchPresenter {
	return LaunchPresenter{
		ID:            launch.LaunchId,
		LaunchDate:    launch.LaunchDate,
		LaunchSuccess: launch.LaunchSuccess,
		VideoLink:     launch.VideoLink,
		MissionName:   launch.MissionName,
		RocketName:    launch.RocketName,
		LaunchDetails: launch.LaunchDetails,
	}
}

// Format - creates a launch representation from the Launch entity
func (p LaunchPresenter) Format() []byte {
	response, err := json.Marshal(p)
	if err != nil {
		return []byte{}
	}
	return response
}

// FormatMany - creates multiple launch representations from a batch of Launches entity
func FormatMany(p []LaunchPresenter) []byte {
	if p == nil {
		return []byte{}
	}
	response, err := json.Marshal(p)
	if err != nil {
		return []byte{}
	}
	return response
}
