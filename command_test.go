package gobot

import (
	"testing"
)

func TestPlaceIsFirstAllowedCommand(t *testing.T) {

	r := getTestRobot()

	cmd := Command{}
	cmd.Command = "MOVE"
	response := r.HandleCommand(cmd)
	if response.Error == nil {
		t.Errorf("No error returned when MOVE before PLACE")
		t.Fail()
	}

	cmd = Command{}
	cmd.Command = "LEFT"
	response = r.HandleCommand(cmd)
	if response.Error == nil {
		t.Errorf("No error returned when LEFT before PLACE")
		t.Fail()
	}

	cmd = Command{}
	cmd.Command = "RIGHT"
	response = r.HandleCommand(cmd)
	if response.Error == nil {
		t.Errorf("No error returned when RIGHT before PLACE")
		t.Fail()
	}

	cmd = Command{}
	cmd.Command = "REPORT"
	response = r.HandleCommand(cmd)
	if response.Error == nil {
		t.Errorf("No error returned when REPORT before PLACE")
		t.Fail()
	}
}
