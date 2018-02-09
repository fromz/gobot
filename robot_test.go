package gobot

import (
	"testing"
)

func getTestRobot() Robot {
	table := Grid{
		Height: 5,
		Width:  5,
	}

	table.Init()

	return Robot{
		Name: "Rudy",
		Icon: "r",
		Grid: table,
		CoordinateConverter: SouthWestConverter{
			table.Height,
			table.Width,
		},
	}

}

func testCommandsExpectedReport(t *testing.T, cmds []string, expectedSpeak string) {
	r := getTestRobot()
	for _, cmdInput := range cmds {
		cmd, _ := Parse(cmdInput)
		r.HandleCommand(cmd)
	}

	cmd := Command{}
	cmd.Command = "REPORT"
	response := r.HandleCommand(cmd)
	if response.Error != nil || response.Speak != expectedSpeak {
		t.Errorf("Expected %s, got %s", expectedSpeak, response.Speak)
	}
}

func TestRobotMoveOne(t *testing.T) {

	expectedSpeak := "0,1,NORTH"
	cmds := []string{
		"PLACE 0,0,NORTH",
		"MOVE",
	}

	testCommandsExpectedReport(t, cmds, expectedSpeak)
}

func TestRobotMoveTwo(t *testing.T) {

	expectedSpeak := "0,0,WEST"
	cmds := []string{
		"PLACE 0,0,NORTH",
		"LEFT",
	}

	testCommandsExpectedReport(t, cmds, expectedSpeak)
}

func TestRobotMoveThree(t *testing.T) {

	expectedSpeak := "3,3,NORTH"
	cmds := []string{
		"PLACE 1,2,EAST",
		"MOVE",
		"MOVE",
		"LEFT",
		"MOVE",
	}

	testCommandsExpectedReport(t, cmds, expectedSpeak)
}
