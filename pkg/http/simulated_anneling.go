package http

import (
	"encoding/json"

	"github.com/patryklyczko/transport_app/pkg/db"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) anneling(ctx *fasthttp.RequestCtx) {
	var order *db.OrderRequest
	var id string
	var err error

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal order %v", err)
		return
	}

	if err = i.api.Anneling(nil); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while adding order %v", err)
		return
	}

	i.log.Debugf("added order %v", id)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}
