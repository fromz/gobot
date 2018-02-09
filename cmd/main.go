package main

import (
	"bufio"
	"fmt"
	"github.com/fromz/gobot"
	"os"
	"strings"
)

func main() {
	table := gobot.Grid{
		Height: 5,
		Width:  5,
	}

	table.Init()

	robot := gobot.Robot{
		Name: "Rudy",
		Icon: "r",
		Grid: table,
		CoordinateConverter: gobot.SouthWestConverter{
			table.Height,
			table.Width,
		},
	}

	for {
		table.RenderCli()

		command, err := waitForCommand()
		if err != nil {
			fmt.Println(err)
			continue
		}

		commandResponse := robot.HandleCommand(command)
		if commandResponse.Error != nil {
			fmt.Println(fmt.Sprintf("%s the gobot says: ERROR %s", robot.Name, commandResponse.Error))
		}
		if len(commandResponse.Speak) > 0 {
			fmt.Println(fmt.Sprintf("%s the gobot says: %s", robot.Name, commandResponse.Speak))
		}
	}
}

func waitForCommand() (gobot.Command, error) {
	c := gobot.Command{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Command: ")
	cin, err := reader.ReadString('\n')
	if err != nil {
		return c, err
	}
	c, err = gobot.Parse(strings.TrimRight(cin, "\n"))
	if err != nil {
		return c, err
	}
	return c, nil
}
