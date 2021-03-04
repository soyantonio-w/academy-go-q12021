package presenter

import (
	"testing"
)

func TestLaunch(t *testing.T) {
	p := LaunchPresenter{
		ID: 2,
		LaunchDate: "",
		LaunchDetails: "",
		LaunchSuccess: false,
		VideoLink: "",
		RocketName: "",
		MissionName: "",
	}
	expect := "{" +
		"\"id\":2," +
		"\"launch_date\":\"\"," +
		"\"success\":false," +
		"\"details\":\"\"," +
		"\"mission_name\":\"\"," +
		"\"rocket_name\":\"\"," +
		"\"video_link\":\"\"" +
		"}"

	if got := string(p.Format()); got != expect{
		t.Errorf("Expect %s, got %s", expect, got)
	}
}
