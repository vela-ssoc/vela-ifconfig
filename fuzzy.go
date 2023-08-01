package ifconfig

func noop(p Interface) bool {
	return true
}

func fuzzy(match func(string) bool) func(Interface) bool {
	return func(i Interface) bool {
		if match(i.Info.Name) {
			return true
		}

		if match(i.Info.HardwareAddr.String()) {
			return true
		}

		if match(i.Info.Flags.String()) {
			return true
		}

		if filterByAddr(i, match) {
			return true
		}

		return false
	}
}

func filterByAddr(ifi Interface, match func(string) bool) bool {
	addr := ifi.Addr()
	n := len(addr)
	if n == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if match(addr[i].IP) {
			return true
		}
	}

	return false
}
