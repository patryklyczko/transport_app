package http

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/patryklyczko/transport_app/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type HTTPInstanceAPI struct {
	bind string
	log  logrus.FieldLogger
	api  *api.InstanceAPI
}

func NewHTTPInstanceAPI(bind string, log logrus.FieldLogger, api *api.InstanceAPI) *HTTPInstanceAPI {
	return &HTTPInstanceAPI{
		bind: bind,
		log:  log,
		api:  api,
	}
}

func (i *HTTPInstanceAPI) Run() {
	r := router.New()

	r.GET("/", i.handleRoot)
	r.POST("/process_map", i.processMap)

	// Order
	r.GET("/orders", i.orders)
	r.GET("/order/{id}", i.order)
	r.GET("/orders/{page}_{number}", i.pageOrder)
	r.POST("/orders", i.addOrders)
	r.PUT("/orders", i.updateOrder)
	r.DELETE("/orders/{id}", i.deleteOrder)

	// Driver
	r.GET("/drivers", i.drivers)
	r.GET("/driver/{id}", i.driver)
	r.GET("/drivers/{page}_{number}", i.pageDriver)
	r.POST("/driver", i.addDrivers)
	r.PUT("/driver", i.updateDriver)
	r.DELETE("/driver/{id}", i.deleteDriver)

	// Algorithms
	r.POST("/simulated_anneling", i.anneling)

	cors := func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
			ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
			h(ctx)
		}
	}

	i.log.Infof("Starting server at %s", i.bind)
	s := &fasthttp.Server{
		Handler:            cors(r.Handler),
		Name:               "Transport_app",
		MaxRequestBodySize: 64 * 1024 * 1024 * 1024, // 64MiB
	}
	log.Fatal(s.ListenAndServe(i.bind))
}

func (i *HTTPInstanceAPI) handleRoot(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetBodyString("Welcome!!")
}
