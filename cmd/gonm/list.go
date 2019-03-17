package cmd

import (
	"os"

	"github.com/nirmoy/gonm/pkg/netutils"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	allDevices bool
)
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list network interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		//allInterfaces := [][]string{}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Flags", "Addr"})
		table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgWhiteColor},
			tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})
		addrs := netutils.GetAddrs()
		for _, addr := range addrs {
			table.Append(addr)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allDevices, "all", "a", false, "List all devices")
}
