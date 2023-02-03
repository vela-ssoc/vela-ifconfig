package ifconfig

import (
	"bytes"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/auxlib"
	"github.com/vela-ssoc/vela-kit/lua"
	"net"
	"strconv"
)

func (ifi *Interface) String() string                         { return "" }
func (ifi *Interface) Type() lua.LValueType                   { return lua.LTObject }
func (ifi *Interface) AssertFloat64() (float64, bool)         { return 0, false }
func (ifi *Interface) AssertString() (string, bool)           { return "", false }
func (ifi *Interface) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (ifi *Interface) Peek() lua.LValue                       { return ifi }

func (ifi *Interface) addrL(L *lua.LState) int {
	n := L.CheckInt(1)

	addr, err := ifi.face.Addrs()
	if err != nil {
		return 0
	}

	if n >= len(addr) {
		return 0
	}
	ip := addr[n].(*net.IPNet)

	L.Push(lua.S2L(ip.IP.String()))
	L.Push(lua.S2L(ip.Mask.String()))
	return 2
}

func (ifi *Interface) helper(match func(string) bool) string {
	addr, err := ifi.face.Addrs()
	if err != nil {
		return ""
	}

	n := len(addr)
	if n == 0 {
		return ""
	}

	var buf bytes.Buffer
	k := 0
	for i := 0; i < n; i++ {
		ip := addr[i].(*net.IPNet).IP.String()

		if !match(ip) {
			continue
		}

		if k > 0 {
			buf.WriteByte(',')
		}
		k++

		buf.WriteString(ip)
	}
	return buf.String()
}

func (ifi *Interface) ipv4L() string {
	return ifi.helper(auxlib.Ipv4)
}

func (ifi *Interface) ipv6L() string {
	return ifi.helper(auxlib.Ipv6)
}

func (ifi *Interface) CompareIP(filter func(string) bool, val string, method cond.Method) bool {
	naddr, err := ifi.face.Addrs()
	if err != nil {
		return false
	}

	for _, addr := range naddr {
		ip := addr.(*net.IPNet).IP.String()
		if filter != nil && !filter(ip) {
			continue
		}

		if method(ip, val) {
			return true
		}
	}
	return false

}

func (ifi *Interface) Compare(key string, val string, method cond.Method) bool {
	switch key {

	case "name":
		return method(ifi.face.Name, val)

	case "flag":
		return method(ifi.face.Flags.String(), val)

	case "index":
		return method(strconv.Itoa(ifi.face.Index), val)

	case "mac":
		return method(ifi.Mac(), val)

	case "mtu":
		return method(strconv.Itoa(ifi.face.MTU), val)

	case "addr":
		return ifi.CompareIP(nil, val, method)

	case "ipv4":
		return ifi.CompareIP(auxlib.Ipv4, val, method)

	case "ipv6":
		return ifi.CompareIP(auxlib.Ipv6, val, method)
	}

	return false

}

func (ifi *Interface) Index(L *lua.LState, key string) lua.LValue {

	switch key {

	case "name":
		return lua.S2L(ifi.face.Name)

	case "flag":
		return lua.S2L(ifi.face.Flags.String())

	case "index":
		return lua.LInt(ifi.face.Index)

	case "mac":
		return lua.S2L(ifi.Mac())

	case "mtu":
		return lua.LInt(ifi.face.MTU)

	case "addr":
		return lua.NewFunction(ifi.addrL)

	case "ipv4":
		return lua.S2L(ifi.ipv4L())

	case "ipv6":
		return lua.S2L(ifi.ipv6L())

	}

	return lua.LNil

}
