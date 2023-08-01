package ifconfig

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/vela-ssoc/vela-kit/vela"
)

func define(r vela.Router) {
	r.GET("/api/v1/arr/agent/interface/list", xEnv.Then(func(ctx *fasthttp.RequestCtx) error {
		_G.update()
		if _G.Err != nil {
			return _G.Err
		}

		chunk, err := json.Marshal(_G.Entry)
		if err != nil {
			return err
		}
		ctx.Write(chunk)
		return nil
	}))
}
