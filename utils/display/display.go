package display

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)

	for i, v := range data {
		if i == 0 {
			table.Header(v)
			continue
		}
		table.Append(v)
	}

	table.Render()
}
