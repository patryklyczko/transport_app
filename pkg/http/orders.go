package http

import (
	"encoding/json"
	"strconv"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) orders(ctx *fasthttp.RequestCtx) {
	var orders []structures.Order
	var err error
	var body []byte

	if orders, err = i.api.Orders(); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting orders: %v", err)
		return
	}

	if body, err = json.Marshal(orders); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) order(ctx *fasthttp.RequestCtx) {
	var order *structures.Order
	var err error
	var body []byte

	id := ctx.UserValue("id").(string)

	if order, err = i.api.Order(id); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting orders: %v", err)
		return
	}

	if body, err = json.Marshal(&order); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) pageOrder(ctx *fasthttp.RequestCtx) {
	var orders *structures.OrderPagination
	var err error
	var body []byte

	page, _ := strconv.Atoi(ctx.UserValue("page").(string))
	number, _ := strconv.Atoi(ctx.UserValue("number").(string))

	if orders, err = i.api.PageOrder(page, number); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting pagination orders: %v", err)
		return
	}

	if body, err = json.Marshal(orders); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) addOrders(ctx *fasthttp.RequestCtx) {
	var order *structures.OrderRequest
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
	var order *structures.Order
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
	id := ctx.UserValue("id").(string)

	if err = i.api.DeleteOrder(id); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while deleting order %v", err)
		return
	}

	i.log.Debugf("deleted order %v", id)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}
