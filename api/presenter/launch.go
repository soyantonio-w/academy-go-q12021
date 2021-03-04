package presenter

import (
	"encoding/json"
	"github.com/soyantonio-w/academy-go-q12021/entity"
)

type LaunchPresenter struct {
	ID entity.LaunchId 				`json:"id"`
	LaunchDate entity.LaunchDate 	`json:"launch_date"`
	entity.LaunchSuccess			`json:"success"`
	entity.LaunchDetails			`json:"details"`
	entity.MissionName				`json:"mission_name"`
	entity.RocketName				`json:"rocket_name"`
	entity.VideoLink				`json:"video_link"`
}

func NewLaunchPresenter(launch entity.Launch) LaunchPresenter {
	return LaunchPresenter{
		ID: launch.LaunchId,
		LaunchDate: launch.LaunchDate,
		LaunchSuccess: launch.LaunchSuccess,
		VideoLink: launch.VideoLink,
		MissionName: launch.MissionName,
		RocketName: launch.RocketName,
		LaunchDetails: launch.LaunchDetails,
	}
}

func (p LaunchPresenter) Format() []byte {
	response, err := json.Marshal(p)
	if err != nil {
		return []byte{}
	}
	return response
}

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