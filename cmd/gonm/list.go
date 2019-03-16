package cmd

import (
	"net"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

var (
	allDevices bool
)
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list smart TVs in the network",
	Run: func(cmd *cobra.Command, args []string) {
		//allInterfaces := [][]string{}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Flags", "Addr"})
		table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgWhiteColor},
			tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})
		ifaces, _ := net.Interfaces()

		for _, iface := range ifaces {
			firstElem := true
			link, _ := netlink.LinkByName(iface.Name)
			addrs, _ := netlink.AddrList(link, netlink.FAMILY_V4)

			for _, addr := range addrs {
				if firstElem {
					firstElem = false
					table.Append([]string{iface.Name, iface.Flags.String(), addr.IPNet.String()})
				} else {
					table.Append([]string{"", "", addr.IPNet.String()})
				}
			}
			addrs, _ = netlink.AddrList(link, netlink.FAMILY_V6)
			for _, addr := range addrs {
				if firstElem {
					firstElem = false
					table.Append([]string{iface.Name, iface.Flags.String(), addr.IPNet.String()})
				} else {
					table.Append([]string{"", "", addr.IPNet.String()})
				}
			}
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allDevices, "all", "a", false, "List all devices")
}
