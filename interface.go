package ifconfig

import (
	"net"
	"time"
)

type Addr struct {
	IP   string `json:"ip"`
	Mask string `json:"mask"`
}
type Interface struct {
	Info  net.Interface `json:"info"`
	Flow  *flow         `json:"flow"`
	Last  time.Time     `json:"last"`
	Addrs []Addr        `json:"addrs"`
}

func (ifi *Interface) patch() {
	addrs, err := ifi.Info.Addrs()
	if err != nil {
		xEnv.Errorf("network interface lookup fail %v", err)
		return
	}

	for _, item := range addrs {
		addr := item.(*net.IPNet)
		ifi.Addrs = append(ifi.Addrs, Addr{
			IP:   addr.IP.String(),
			Mask: addr.Mask.String(),
		})
	}
}

func (ifi *Interface) Name() string {
	return ifi.Info.Name
}

func (ifi *Interface) Addr() []Addr {
	return ifi.Addrs
}

func (ifi *Interface) Mac() string {
	if ifi.Info.HardwareAddr == nil {
		return ""
	}

	return ifi.Info.HardwareAddr.String()
}
