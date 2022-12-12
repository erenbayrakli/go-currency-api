package main

import (
	"currency-api/services"
	"log"
	"strings"

	"github.com/fasthttp/router"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	mark := strings.Trim(string(ctx.Request.RequestURI()), "/")
	uri := strings.ToUpper(mark)
	if uri == "" {
		res := services.GetCurrencies()
		ctx.SetStatusCode(200)
		ss, _ := jsoniter.MarshalToString(res)
		ctx.Write([]byte(ss))
	} else {
		res := services.GetCurrenyByName(uri).ForexSelling
		ctx.Response.Header.Set("content-type", "application/json")
		ctx.SetStatusCode(200)
		ctx.Write([]byte(res))
	}
}
func main() {
	r := router.New()
	r.GET("/", requestHandler)
	r.GET("/usd", requestHandler)
	r.GET("/eur", requestHandler)
	r.GET("/cad", requestHandler)
	r.GET("/aud", requestHandler)
	r.GET("/gbp", requestHandler)
	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))

}
