package entity

import "testing"

func TestNewLaunch(t *testing.T) {
	expectedId := 21
	expectedDetails := "details..."
	expectedDate := "2020-02-12"
	expectedRocket := "Falcon 9"
	expectedMission := "Base"
	expectedLink := "http://dd.com"
	expectedSuccess := false
	got := NewLaunch(expectedId, expectedDetails, expectedDate, expectedMission, expectedLink, expectedRocket, expectedSuccess)

	if int(got.LaunchId) != expectedId {
		t.Errorf("Expected %d, got %d", expectedId, int(got.LaunchId))
	}
	expectString(t, string(got.LaunchDetails), expectedDetails)
	expectString(t, string(got.LaunchDate), expectedDate)
	expectString(t, string(got.RocketName), expectedRocket)
	expectString(t, string(got.MissionName), expectedMission)
	expectString(t, string(got.VideoLink), expectedLink)

	if bool(got.LaunchSuccess) != expectedSuccess {
		t.Errorf("Expected %v, got %v", expectedSuccess, bool(got.LaunchSuccess))
	}
}

func expectString(t *testing.T, got, expect string) {
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}
