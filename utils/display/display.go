package display

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(data[1:])
	table.Render()
}
