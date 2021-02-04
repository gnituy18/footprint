package main

import (
	"github.com/valyala/fasthttp"

	"footprint/pkg/api"
)

func main() {
	fasthttp.ListenAndServe(":8080", api.Router().HandleRequest)
}
