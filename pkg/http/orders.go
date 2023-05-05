package http

import (
	"encoding/json"

	"github.com/patryklyczko/transport_app/pkg/db"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) addOrders(ctx *fasthttp.RequestCtx) {
	var order *db.OrderRequest
	var id string
	var err error

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal order %v", err)
		return
	}

	if id, err = i.api.AddOrder(order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while adding order %v", err)
		return
	}

	i.log.Debugf("added order %v", id)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) updateOrder(ctx *fasthttp.RequestCtx) {
	var order *db.Order
	var err error

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal order %v", err)
		return
	}

	if err = i.api.UpdateOrder(order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while adding order %v", err)
		return
	}

	i.log.Debugf("updated order %v", order.ID)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) deleteOrder(ctx *fasthttp.RequestCtx) {
	var err error
	ID := ctx.UserValue("uid").(string)
	if err = i.api.DeleteOrder(ID); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while deleting order %v", err)
		return
	}

	i.log.Debugf("deleted order %v", ID)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}
