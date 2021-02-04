package main

import (
	"github.com/valyala/fasthttp"

	"footprint/pkg/api"
	"footprint/pkg/log"
)

func main() {
	if err := fasthttp.ListenAndServe(":8080", api.Router().HandleRequest); err != nil {
		log.Global().Error("fasthttp.ListenAndServe failed in main")
	}
}
