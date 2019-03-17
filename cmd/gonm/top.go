package cmd

import (
	"fmt"
	"log"
	"net"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

func initIfaceGauge(ifName string, rect []int) (*widgets.SparklineGroup, error) {
	sl := widgets.NewSparkline()
	sl.Data = []float64{}
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	sl2 := widgets.NewSparkline()
	sl2.Data = []float64{}
	sl2.TitleStyle.Fg = ui.ColorWhite
	sl2.LineColor = ui.ColorRed

	slg := widgets.NewSparklineGroup(sl, sl2)
	slg.Title = ifName
	slg.SetRect(rect[0], rect[1], rect[2], rect[3])

	go func() {
		link, _ := netlink.LinkByName(ifName)
		prevRxBytes := link.Attrs().Statistics.RxBytes
		prevTxBytes := link.Attrs().Statistics.TxBytes
		for {
			link, _ := netlink.LinkByName(ifName)
			currRxBytes := link.Attrs().Statistics.RxBytes
			currTxBytes := link.Attrs().Statistics.TxBytes
			rxBytes := currRxBytes - prevRxBytes
			txBytes := currTxBytes - prevTxBytes

			sl.Data = append(sl.Data, float64(rxBytes))
			sl2.Data = append(sl2.Data, float64(txBytes))

			slg.Sparklines[0].Title = fmt.Sprintf(" Rx throughput: %v %v", rxBytes, "Bytes/Sec")
			slg.Sparklines[1].Title = fmt.Sprintf(" Tx throughput: %v %v", txBytes, "Bytes/Sec")
			time.Sleep(1 * time.Second)
			prevRxBytes = link.Attrs().Statistics.RxBytes
			prevTxBytes = link.Attrs().Statistics.TxBytes

		}

	}()
	return slg, nil
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "live stats of all interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		var allGauge []ui.Drawable
		var allPane []string
		var index int

		ifaces, _ := net.Interfaces()

		for _, iface := range ifaces {
			wg, _ := initIfaceGauge(iface.Name, []int{0, 5, 80, 15})

			allGauge = append(allGauge, wg)
			allPane = append(allPane, iface.Name)
			index++
		}

		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		header := widgets.NewParagraph()
		header.Text = "Press q to quit, Press h or l to switch tabs"
		header.SetRect(0, 0, 50, 1)
		header.Border = false
		header.TextStyle.Bg = ui.ColorBlue

		tabpane := widgets.NewTabPane(allPane...)
		tabpane.SetRect(0, 1, 80, 4)
		tabpane.Border = true
		ui.Render(tabpane)
		uiEvents := ui.PollEvents()
		ticker := time.NewTicker(time.Second).C
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "h":
					tabpane.FocusLeft()
					ui.Clear()
					ui.Render(header, tabpane)
					ui.Render(header, allGauge[tabpane.ActiveTabIndex])
				case "l":
					tabpane.FocusRight()
					ui.Clear()
					ui.Render(header, tabpane)
					ui.Render(header, allGauge[tabpane.ActiveTabIndex])
				}
			case <-ticker:
				ui.Render(header, tabpane)
				ui.Render(allGauge[tabpane.ActiveTabIndex])
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
