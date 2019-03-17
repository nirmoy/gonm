package cmd

import (
	"log"
	"net"
	"time"

	gauge "github.com/nirmoy/gonm/pkg/widgets"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

var (
	sparkLingLen = 80
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "live stats of all interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		var allGauge []ui.Drawable
		var allPane []string

		ifaces, _ := net.Interfaces()

		for _, iface := range ifaces {
			wg, _ := gauge.InitIfaceGauge(iface.Name, []int{0, 5, sparkLingLen, 15})

			allGauge = append(allGauge, wg)
			allPane = append(allPane, iface.Name)
		}

		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		header := widgets.NewParagraph()
		header.Text = "Press q to quit, Press h or l to switch tabs"
		header.SetRect(0, 0, sparkLingLen, 1)
		header.Border = false
		header.TextStyle.Bg = ui.ColorBlue

		tabpane := widgets.NewTabPane(allPane...)
		tabpane.SetRect(0, 1, sparkLingLen, 4)
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
