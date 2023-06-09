package http

import (
	"encoding/json"

	"github.com/patryklyczko/transport_app/pkg/db"
	"github.com/patryklyczko/transport_app/pkg/structures"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) anneling(ctx *fasthttp.RequestCtx) {
	var parameters *db.AnnelingParameters
	var solutionValue *structures.SolutionValues
	var err error
	// var solution map[*db.Driver][]db.Order
	// var gain float32

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &parameters); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while algorithm anneling %v", err)
		return
	}

	if solutionValue, err = i.api.Anneling(parameters); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while algorithm anneling %v", err)
		return
	}

	if body, err = json.Marshal(solutionValue); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}
	ctx.Response.SetStatusCode(fasthttp.StatusAccepted)
	ctx.Response.SetBody(body)
}
