package netutils

import (
	"net"

	"github.com/vishvananda/netlink"
)

func GetAddrs() [][]string {
	var result [][]string
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		firstElem := true
		link, _ := netlink.LinkByName(iface.Name)
		addrs, _ := netlink.AddrList(link, netlink.FAMILY_V4)

		for _, addr := range addrs {
			if firstElem {
				firstElem = false
				result = append(result, []string{iface.Name, iface.Flags.String(), addr.IPNet.String()})
			} else {
				result = append(result, []string{"", "", addr.IPNet.String()})
			}
		}
		addrs, _ = netlink.AddrList(link, netlink.FAMILY_V6)
		for _, addr := range addrs {
			if firstElem {
				firstElem = false
				result = append(result, []string{iface.Name, iface.Flags.String(), addr.IPNet.String()})
			} else {
				result = append(result, []string{"", "", addr.IPNet.String()})
			}
		}
	}
	return result
}
