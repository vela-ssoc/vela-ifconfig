package ifconfig

import (
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/kind"
	"github.com/vela-ssoc/vela-kit/lua"
	"net"
	"time"
)

type summary struct {
	Entry []Interface
	Err   error
}

func (sum *summary) Byte() []byte {
	enc := kind.NewJsonEncoder()
	enc.Arr("")
	for _, v := range sum.Entry {
		enc.Tab("")
		enc.KV("index", v.Info.Index)
		enc.KV("mtu", v.Info.MTU)
		enc.KV("name", v.Info.Name)
		enc.KV("mac", v.Info.HardwareAddr.String())
		enc.KV("flags", v.Info.Flags.String())
		//enc.Raw("flow", v.flow.Byte())
		enc.End("},")
	}
	enc.End("]")
	return enc.Bytes()
}

func (sum *summary) Len() int {
	return len(sum.Entry)
}

func (sum *summary) append(vi Interface) {
	sum.Entry = append(sum.Entry, vi)
}

func (sum *summary) update() {
	face, err := net.Interfaces()
	if err != nil {
		sum.Err = err
		return
	}

	n := len(face)
	now := time.Now()

	entry := make([]Interface, n)
	for i := 0; i < n; i++ {
		ifc := Interface{Info: face[i], Last: now}
		ifc.patch()
		entry[i] = ifc
	}

	sum.Entry = entry
	sum.flow()
}

func (sum *summary) ok() bool {
	if sum.Err != nil {
		return false
	}

	return true
}

func (sum *summary) lookup(cnd *cond.Cond) (ret lua.Slice) {
	if sum.Err != nil {
		return
	}

	n := sum.Len()
	for i := 0; i < n; i++ {
		ifi := sum.Entry[i]
		if cnd != nil && !cnd.Match(&ifi) {
			continue
		}
		ret = append(ret, &ifi)
	}
	return
}
