package http

import (
	"encoding/json"
	"strconv"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) drivers(ctx *fasthttp.RequestCtx) {
	var drivers []structures.Driver
	var err error
	var body []byte

	if drivers, err = i.api.Drivers(); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting drivers: %v", err)
		return
	}

	if body, err = json.Marshal(drivers); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) driver(ctx *fasthttp.RequestCtx) {
	var driver *structures.Driver
	var err error
	var body []byte

	id := ctx.UserValue("id").(string)

	if driver, err = i.api.Driver(id); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting driver: %v", err)
		return
	}

	if body, err = json.Marshal(&driver); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) pageDriver(ctx *fasthttp.RequestCtx) {
	var orders *structures.DriverPagination
	var err error
	var body []byte

	page, _ := strconv.Atoi(ctx.UserValue("page").(string))
	number, _ := strconv.Atoi(ctx.UserValue("number").(string))

	if orders, err = i.api.PageDriver(page, number); err != nil {
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

func (i *HTTPInstanceAPI) addDrivers(ctx *fasthttp.RequestCtx) {
	var driver *structures.DriverRequest
	var id string
	var err error

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &driver); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal driver %v", err)
		return
	}

	if id, err = i.api.AddDriver(driver); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while adding driver %v", err)
		return
	}

	i.log.Debugf("added driver %v", id)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) updateDriver(ctx *fasthttp.RequestCtx) {
	var driver *structures.Driver
	var err error

	body := ctx.Request.Body()
	if err = json.Unmarshal(body, &driver); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshal driver %v", err)
		return
	}

	if err = i.api.UpdateDriver(driver); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while adding driver %v", err)
		return
	}

	i.log.Debugf("updated driver %v", driver.ID)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) deleteDriver(ctx *fasthttp.RequestCtx) {
	var err error
	id := ctx.UserValue("id").(string)

	if err = i.api.DeleteDriver(id); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while deleting driver %v", err)
		return
	}

	i.log.Debugf("deleted driver %v", id)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}
