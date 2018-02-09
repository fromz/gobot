package gobot

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// RenderCli renders the grid on the command line
func (g *Grid) RenderCli() {
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 0, '.', tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent)
	for _, row := range g.Objects {
		display := []string{}
		for _, gridObject := range row {
			display = append(display, gridObject.DecorateSquare())
		}
		fmt.Fprintln(w, strings.Join(display, "\t")+"\t")
	}
	w.Flush()
}
