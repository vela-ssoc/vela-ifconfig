package ifconfig

import (
	"github.com/vela-ssoc/vela-kit/lua"
)

type flow struct {
	InBytes         uint64  `json:"in_bytes"`
	InPackets       uint64  `json:"in_packets"`
	InError         uint64  `json:"in_error"`
	InDropped       uint64  `json:"in_dropped"`
	InBytesPerSec   float64 `json:"in_bytes_per_sec"`
	InPacketsPerSec float64 `json:"in_packets_per_sec"`

	OutBytes         uint64  `json:"out_bytes"`
	OutPackets       uint64  `json:"out_packets"`
	OutError         uint64  `json:"out_error"`
	OutDropped       uint64  `json:"out_dropped"`
	OutBytesPerSec   float64 `json:"out_bytes_per_sec"`
	OutPacketsPerSec float64 `json:"out_packets_per_sec"`
}

func (f *flow) ToLValue() lua.LValue {
	return lua.NewAnyData(f)
}

func (f *flow) clone() *flow {
	return &flow{
		InBytes:          f.InBytes,
		InPackets:        f.InPackets,
		InError:          f.InError,
		InDropped:        f.InDropped,
		InBytesPerSec:    f.InBytesPerSec,
		InPacketsPerSec:  f.InPacketsPerSec,
		OutBytes:         f.OutBytes,
		OutPackets:       f.OutPackets,
		OutError:         f.OutError,
		OutDropped:       f.OutDropped,
		OutBytesPerSec:   f.OutBytesPerSec,
		OutPacketsPerSec: f.OutPacketsPerSec,
	}
}
