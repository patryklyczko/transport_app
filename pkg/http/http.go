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

func (i *HTTPInstanceAPI) OptionsHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "http://localhost:3000")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Expose-Headers", "*")
	ctx.Response.SetStatusCode(200)
}

func (i *HTTPInstanceAPI) Run() {
	r := router.New()

	r.GET("/", i.handleRoot)
	// Put manual video
	r.POST("/process_map", i.processMap)

	// Order
	r.GET("/orders", i.orders)
	r.GET("/order", i.order)
	r.POST("/order", i.addOrders)
	r.PUT("/order", i.updateOrder)
	r.DELETE("/order", i.deleteOrder)

	// Driver
	// r.POST("/driver")

	i.log.Infof("Starting server at %s", i.bind)
	s := &fasthttp.Server{
		Handler:            r.Handler,
		Name:               "Video_Streaming",
		MaxRequestBodySize: 64 * 1024 * 1024 * 1024, // 64MiB
	}
	log.Fatal(s.ListenAndServe(i.bind))
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Expose-Headers", "*")
	ctx.Response.Header.Add("Content-type", "application/json charset=utf-8 video/mp4")
}

func (i *HTTPInstanceAPI) handleRoot(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetBodyString("Welcome!!")
}
