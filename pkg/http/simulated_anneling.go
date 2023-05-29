package http

import (
	"encoding/json"

	"github.com/patryklyczko/transport_app/pkg/db"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) anneling(ctx *fasthttp.RequestCtx) {
	var parameters *db.AnnelingParameters
	var err error
	// var solution map[*db.Driver][]db.Order
	// var gain float32

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &parameters); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while algorithm anneling %v", err)
		return
	}

	if _, _, err = i.api.Anneling(parameters); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while algorithm anneling %v", err)
		return
	}
	// i.log.Debugf("Solution %v \n gain %v", solution, gain)

	ctx.Response.SetStatusCode(fasthttp.StatusAccepted)
}
