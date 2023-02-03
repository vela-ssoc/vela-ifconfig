package ifconfig

import (
	"github.com/bytedance/sonic"
	"github.com/vela-ssoc/vela-kit/lua"
)

func (f *flow) String() string                         { return lua.B2S(f.Byte()) }
func (f *flow) Type() lua.LValueType                   { return lua.LTObject }
func (f *flow) AssertFloat64() (float64, bool)         { return 0, false }
func (f *flow) AssertString() (string, bool)           { return "", false }
func (f *flow) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (f *flow) Peek() lua.LValue                       { return f }

func (f *flow) Byte() []byte {
	chunk, _ := sonic.Marshal(f)
	return chunk
}

func (f *flow) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "in_bytes":
		return lua.LNumber(f.InBytes)
	case "in_packets":
		return lua.LNumber(f.InPackets)
	case "in_error":
		return lua.LNumber(f.InError)
	case "in_dropped":
		return lua.LNumber(f.InDropped)
	case "in_bps":
		return lua.LNumber(f.InBytesPerSec)
	case "in_pps":
		return lua.LNumber(f.InPacketsPerSec)

	case "out_bytes":
		return lua.LNumber(f.OutBytes)
	case "out_packets":
		return lua.LNumber(f.OutPackets)
	case "out_error":
		return lua.LNumber(f.OutError)
	case "out_dropped":
		return lua.LNumber(f.OutDropped)
	case "out_bps":
		return lua.LNumber(f.OutBytesPerSec)
	case "out_pps":
		return lua.LNumber(f.OutPacketsPerSec)
	}
	return lua.LNil
}
