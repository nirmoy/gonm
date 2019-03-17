package widgets

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/vishvananda/netlink"
)

func InitIfaceGauge(ifName string, rect []int) (*widgets.SparklineGroup, error) {
	sparkLingLen := rect[2]

	sl1 := widgets.NewSparkline()
	sl1.Data = []float64{}
	sl1.LineColor = ui.ColorCyan
	sl1.TitleStyle.Fg = ui.ColorWhite

	sl2 := widgets.NewSparkline()
	sl2.Data = []float64{}
	sl2.TitleStyle.Fg = ui.ColorWhite
	sl2.LineColor = ui.ColorRed

	slg := widgets.NewSparklineGroup(sl1, sl2)
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

			sl1.Data = append(sl1.Data, float64(rxBytes))
			sl2.Data = append(sl2.Data, float64(txBytes))
			if len(sl1.Data) > sparkLingLen {
				sl1.Data = sl1.Data[1:]
				sl2.Data = sl2.Data[1:]
			}

			slg.Sparklines[0].Title = fmt.Sprintf(" Rx throughput: %v %v", rxBytes, "Bytes/Sec")
			slg.Sparklines[1].Title = fmt.Sprintf(" Tx throughput: %v %v", txBytes, "Bytes/Sec")
			time.Sleep(1 * time.Second)
			prevRxBytes = link.Attrs().Statistics.RxBytes
			prevTxBytes = link.Attrs().Statistics.TxBytes

		}

	}()
	return slg, nil
}
