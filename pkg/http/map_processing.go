package http

import (
	"encoding/json"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) processMap(ctx *fasthttp.RequestCtx) {
	var pathOpen structures.MapRequest
	body := ctx.Request.Body()

	if err := json.Unmarshal(body, &pathOpen); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal map path %v", err)
		return
	}

	if err := i.api.ProcessMap(&pathOpen); err != nil {
		i.log.Debugf("errpr processing map: %v", err)
		return
	}
	i.log.Debugf("processed map %v", pathOpen.Path)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}
