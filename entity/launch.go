package entity

type LaunchId int
type LaunchDetails string
type LaunchDate string
type LaunchSuccess bool

type Launch struct {
	LaunchId
	LaunchDetails
	LaunchDate
	LaunchSuccess
	MissionName
	VideoLink
	RocketName
}

func ToLaunchId(id int) LaunchId {
	var algo LaunchId = LaunchId(id)
	return algo
}

type LaunchRepo interface {
	GetLaunches() ([]Launch, error)
}

func NewLaunch(launchId int, details, date, missionName, videoLink, rocketName string, launchSuccess bool) (l Launch) {
	l = Launch{
		LaunchId: LaunchId(launchId),
		LaunchDetails: LaunchDetails(details),
		LaunchDate: LaunchDate(date),
		LaunchSuccess: LaunchSuccess(launchSuccess),
		VideoLink: VideoLink(videoLink),
		RocketName: RocketName(rocketName),
		MissionName: MissionName(missionName),
	}
	return
}