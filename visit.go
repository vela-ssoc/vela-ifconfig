package ifconfig

import "net"

func IsLoopBack(addr net.Addr) bool {
	ipnet, ok := addr.(*net.IPNet)
	if !ok {
		return false
	}

	return ipnet.IP.IsLoopback()
}

func NotLookBack() ([]net.Addr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ret []net.Addr

	for _, addr := range addrs {
		if !IsLoopBack(addr) {
			ret = append(ret, addr)
		}
	}

	return ret, nil
}

func All() ([]net.Addr, error) {
	return net.InterfaceAddrs()
}

func Have(v string) bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		xEnv.Error("ifconfig got interface addr fail %v", err)
		return false
	}

	for _, addr := range addrs {
		if v == addr.String() {
			return true
		}
	}

	return false
}
