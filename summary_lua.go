package ifconfig

import (
	"fmt"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/auxlib"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
	"github.com/vela-ssoc/vela-kit/worker"
	"gopkg.in/tomb.v2"
	"strings"
	"time"
)

func (sum *summary) String() string                         { return fmt.Sprintf("%p", sum) }
func (sum *summary) Type() lua.LValueType                   { return lua.LTObject }
func (sum *summary) AssertFloat64() (float64, bool)         { return 0, false }
func (sum *summary) AssertString() (string, bool)           { return "", false }
func (sum *summary) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (sum *summary) Peek() lua.LValue                       { return sum }

func (sum *summary) Meta(L *lua.LState, key lua.LValue) lua.LValue {
	switch key.Type() {

	case lua.LTInt:
		return sum.r(int(key.(lua.LInt)))

	case lua.LTNumber:
		return sum.r(int(key.(lua.LNumber)))

	case lua.LTString:
		return sum.Index(L, key.String())
	}

	return lua.LNil
}

func (sum *summary) r(idx int) lua.LValue {
	n := len(sum.Entry)
	if n == 0 {
		return lua.LNil
	}

	if idx-1 > n || idx-1 < 0 {
		return lua.LNil
	}

	return &sum.Entry[idx-1]
}

func (sum *summary) pipeL(L *lua.LState) int {
	pp := pipe.NewByLua(L, pipe.Env(xEnv))

	n := sum.Len()
	if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		vx := sum.Entry[i]
		pp.Do(&vx, L, func(err error) {
			xEnv.Errorf("rock interface pipe fail %v", err)
		})
	}
	return 0
}

func (sum *summary) updateL(L *lua.LState) int {
	sum.update()
	return 0
}

func (sum *summary) toCnd(field, method string, val []string) *cond.Cond {
	if len(val) == 0 {
		return &cond.Cond{}
	}

	return cond.New(field + " " + method + " " + strings.Join(val, ","))
}

func (sum *summary) macL(L *lua.LState) int {
	cnd := sum.toCnd("mac", "=", auxlib.LToSS(L))
	L.Push(sum.lookup(cnd))
	return 1
}

func (sum *summary) addrL(L *lua.LState) int {
	cnd := sum.toCnd("addr", "=", auxlib.LToSS(L))
	L.Push(sum.lookup(cnd))
	return 1
}

func (sum *summary) nameL(L *lua.LState) int {
	cnd := sum.toCnd("name", "=", auxlib.LToSS(L))
	L.Push(sum.lookup(cnd))
	return 1
}

func (sum *summary) flowL(L *lua.LState) int {
	tt := L.IsInt(1)
	//pp := pipe.NewByLua(L, pipe.Seek(1))
	//co := xEnv.Clone(L)

	if tt <= 0 {
		tt = 1000
	}

	tom := new(tomb.Tomb)
	task := func() {
		tk := time.NewTicker(time.Duration(tt) * time.Millisecond)
		defer tk.Stop()

		for {
			select {

			case <-tom.Dying():
				xEnv.Error("rock.interface.flow thread exit")
				return

			case <-tk.C:
				sum.flow()
			}
		}
	}

	kill := func() {
		tom.Kill(fmt.Errorf("over"))
	}

	w := worker.New(L, "rock.interface.flow")
	w.Task(task).Kill(kill).Start()
	return 0
}

func (sum *summary) call(L *lua.LState) int {
	cnd := cond.CheckMany(L)
	L.Push(sum.lookup(cnd))
	return 1
}

func (sum *summary) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	case "update":
		return L.NewFunction(sum.updateL)

	case "pipe":
		return L.NewFunction(sum.pipeL)

	case "addr":
		return L.NewFunction(sum.addrL)

	case "name":
		return L.NewFunction(sum.nameL)

	case "mac":
		return L.NewFunction(sum.macL)

	case "flow":
		return L.NewFunction(sum.flowL)
	}

	return lua.LNil
}
