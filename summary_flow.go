package ifconfig

import (
	"github.com/shirou/gopsutil/net"
	"time"
)

func (sum *summary) flow() {
	sv, e := net.IOCounters(true)
	if e != nil {
		xEnv.Errorf("interface state got fail %v", e)
		return
	}

	if sum.Len() != len(sv) {
		sum.update()
	}

	now := time.Now()
	for i, s := range sv {
		ifi := sum.Entry[i]
		ifi.Last = now
		sum.Entry[i].Flow = &flow{
			InBytes:          s.BytesRecv,
			InPackets:        s.PacketsRecv,
			InError:          s.Errin,
			InDropped:        s.Dropin,
			InBytesPerSec:    0,
			InPacketsPerSec:  0,
			OutBytes:         s.BytesSent,
			OutPackets:       s.PacketsSent,
			OutError:         s.Errout,
			OutDropped:       s.Dropout,
			OutBytesPerSec:   0,
			OutPacketsPerSec: 0,
		}
	}
}
