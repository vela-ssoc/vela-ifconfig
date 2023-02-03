package ifconfig

import (
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/lua"
	"net"
	"time"
)

type summary struct {
	iFace []Interface
	Err   error
}

func (sum *summary) Len() int {
	return len(sum.iFace)
}

func (sum *summary) append(vi Interface) {
	sum.iFace = append(sum.iFace, vi)
}

func (sum *summary) update() {
	face, err := net.Interfaces()
	if err != nil {
		sum.Err = err
		return
	}

	n := len(face)
	now := time.Now()

	for i := 0; i < n; i++ {
		ifc := Interface{face: face[i], last: now}
		sum.iFace = append(sum.iFace, ifc)
	}
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
		ifi := sum.iFace[i]
		if cnd.Match(&ifi) {
			ret = append(ret, &ifi)
		}
	}
	return
}
