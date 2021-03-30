package gspacex

import (
	"fmt"
	"strconv"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

type repository struct {
	url string
}

func NewRepository(url string) entity.LaunchRepo {
	return &repository{
		url: url,
	}
}

type linksParser struct {
	VideoLink string `json:"video_link"`
}

type rocketParser struct {
	Name string `json:"rocket_name"`
}

type launchParser struct {
	Id          string       `json:"id"`
	Details     string       `json:"details"`
	Success     bool         `json:"launch_success"`
	Date        string       `json:"launch_date_utc"`
	MissionName string       `json:"mission_name"`
	Links       linksParser  `json:"links"`
	Rocket      rocketParser `json:"rocket"`
}

type queryLaunchParser struct {
	Data struct {
		// It is a pointer to track when the API does not have a launch
		// https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/
		Launch *launchParser `json:"launch"`
	} `json:"data"`
}

type queryLaunchesParser struct {
	Data struct {
		Launches []launchParser `json:"launches"`
	} `json:"data"`
}

type graphqlVarsLaunch struct {
	Id entity.LaunchId `json:"id"`
}

type graphqlVarsLaunches struct {
	Limit int `json:"limit"`
}

func (repo *repository) Get(id entity.LaunchId) (entity.Launch, error) {
	query := `
	query($id: ID!) {
	  launch(id: $id) {
		id
    	details
		mission_name
		launch_success
		launch_date_utc
		links {
		  video_link
		}
		rocket {
		  rocket_name
		}
	  }
	}
	`
	client := GraphqlClient{endpoint: repo.url}
	response := queryLaunchParser{}
	err := client.load(&response, GraphqlQuery{
		Query: query,
		Variables: graphqlVarsLaunch{
			Id: id,
		},
	})

	if err != nil {
		return entity.Launch{}, err
	}

	launchResponse := response.Data.Launch
	if launchResponse == nil {
		return entity.Launch{}, fmt.Errorf("non found launch with id %d", id)
	}

	return launchResponse.NewLaunch(), nil
}

func (repo *repository) GetLaunches() ([]entity.Launch, error) {
	query := `
		query ($limit: Int) {
		  launches(limit: $limit){
			id
			details
			mission_name
			launch_success
			launch_date_utc
			links {
			  video_link
			}
			rocket {
			  rocket_name
			}
		  }
		}
	`

	client := GraphqlClient{endpoint: repo.url}
	response := queryLaunchesParser{}
	err := client.load(&response, GraphqlQuery{
		Query: query,
		Variables: graphqlVarsLaunches{
			Limit: 100,
		},
	})

	if err != nil {
		return nil, err
	}
	launchesResponse := response.Data.Launches
	if launchesResponse == nil {
		return nil, fmt.Errorf("non launches found")
	}

	var launches []entity.Launch

	for _, launch := range launchesResponse {
		launches = append(launches, launch.NewLaunch())
	}

	return launches, nil
}

func (repo *repository) SyncAll(launches []entity.Launch) error {
	return nil
}

func (lp *launchParser) NewLaunch() (l entity.Launch) {
	launchId, _ := strconv.Atoi(lp.Id)
	return entity.NewLaunch(
		launchId,
		lp.Details,
		lp.Date,
		lp.MissionName,
		lp.Links.VideoLink,
		lp.Rocket.Name,
		lp.Success,
	)
}
