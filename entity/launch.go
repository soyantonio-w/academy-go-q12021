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
