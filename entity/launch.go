package entity

// LaunchId - the id of a Launch
type LaunchId int

// LaunchDetails - description of a Launch
type LaunchDetails string

// LaunchDate - provides the date the Launch occurs
type LaunchDate string

// LaunchSuccess - used to determine a launch mission has been successful
type LaunchSuccess bool

// Launch - object to encapsulate business rules of a launch
type Launch struct {
	LaunchId
	LaunchDetails
	LaunchDate
	LaunchSuccess
	MissionName
	VideoLink
	RocketName
}

// LaunchRepo - use the repo to obtain the information of the launches
type LaunchRepo interface {
	GetLaunches() ([]Launch, error)
	Get(id LaunchId) (Launch, error)
	SyncAll(launches []Launch) error
}

// NewLaunch - creates an instance of Launch
func NewLaunch(launchId int, details, date, missionName, videoLink, rocketName string, launchSuccess bool) (l Launch) {
	l = Launch{
		LaunchId:      LaunchId(launchId),
		LaunchDetails: LaunchDetails(details),
		LaunchDate:    LaunchDate(date),
		LaunchSuccess: LaunchSuccess(launchSuccess),
		VideoLink:     VideoLink(videoLink),
		RocketName:    RocketName(rocketName),
		MissionName:   MissionName(missionName),
	}
	return
}
